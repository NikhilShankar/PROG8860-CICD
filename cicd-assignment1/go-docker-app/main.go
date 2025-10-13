package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
    "io/ioutil"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/time", timeHandler)
    http.HandleFunc("/visit", visitHandler)

    fmt.Printf("Server starting on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Go Docker App</title>
        <style>
            body { font-family: Arial; max-width: 800px; margin: 50px auto; padding: 20px; }
            h1 { color: #00ADD8; }
            .info { background: #f0f0f0; padding: 15px; border-radius: 5px; margin: 10px 0; }
        </style>
    </head>
    <body>
        <h1>🐳 Containerized Go Web Application</h1>
        <div class="info">
            <h2>Welcome!</h2>
            <p>This is a simple Go web application running in a Docker container.</p>
            <h3>Available Endpoints:</h3>
            <ul>
                <li><a href="/">/</a> - Home page (this page)</li>
                <li><a href="/health">/health</a> - Health check endpoint</li>
                <li><a href="/time">/time</a> - Current server time</li>
                <li><a href="/visit">/visit</a> - Record visit (demonstrates volume persistence)</li>
            </ul>
        </div>
    </body>
    </html>
    `
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, html)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, `{"status": "healthy", "service": "go-docker-app"}`)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format("2006-01-02 15:04:05")
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, "<h1>Current Server Time</h1><p>%s</p><a href='/'>Back to Home</a>", currentTime)
}

func visitHandler(w http.ResponseWriter, r *http.Request) {
    logFile := "/app/data/visits.log"
    
    // Create data directory if it doesn't exist
    os.MkdirAll("/app/data", 0755)
    
    // Read existing visits
    content, _ := ioutil.ReadFile(logFile)
    
    // Append new visit
    visit := fmt.Sprintf("Visit at: %s\n", time.Now().Format("2006-01-02 15:04:05"))
    newContent := string(content) + visit
    
    // Write to file
    ioutil.WriteFile(logFile, []byte(newContent), 0644)
    
    w.Header().Set("Content-Type", "text/html")
    html := fmt.Sprintf(`
        <h1>Visit Recorded!</h1>
        <p>Your visit has been logged to persistent storage.</p>
        <h2>All Visits:</h2>
        <pre>%s</pre>
        <a href='/'>Back to Home</a>
    `, newContent)
    fmt.Fprint(w, html)
}