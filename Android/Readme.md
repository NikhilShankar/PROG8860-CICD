



# Counter App - CI/CD Pipeline

## Project Structure

```
repo/
├── android/                    # Android counter app
│   ├── app/
│   │   └── src/
│   │       ├── main/
│   │       │   └── java/rebirth/nixaclabs/cicd/
│   │       │       ├── Counter.kt
│   │       │       └── MainActivity.kt
│   │       └── test/
│   │           └── java/rebirth/nixaclabs/cicd/
│   │               └── CounterTest.kt
│   ├── Dockerfile              # Multi-stage build
│   ├── server.py               # Report server
│   └── build.gradle
├── .github/
│   └── workflows/
│       └── android-ci.yml      # CI Pipeline
└── README.md
```

## App Description

Simple counter app with:
- Display showing current count (starts at 0)
- "+1" button to increment
- "-1" button to decrement

## Build and Run Instructions

### Local Build

```bash
cd android
./gradlew assembleDebug
```

APK location: `app/build/outputs/apk/debug/app-debug.apk`

### Run Tests Locally

```bash
cd android
./gradlew test
```

Test reports: `app/build/reports/tests/testDebugUnitTest/index.html`

### Run Lint Locally

```bash
cd android
./gradlew lint
```

Lint report: `app/build/reports/lint-results-debug.html`

## Docker Instructions

### Pull the Image

```bash
docker pull niks1267/counter-app-reports:latest
```

### Run the Container

```bash
docker run -p 9898:9898 niks1267/counter-app-reports:latest
```

### Access Reports

Open your browser and navigate to:
- **Home**: http://localhost:9898
- **Test Results**: http://localhost:9898/testresults/index.html
- **Lint Report**: http://localhost:9898/lint/index.html

## CI Pipeline

The GitHub Actions workflow:

1. **Build Stage**: Compiles the Android app
2. **Test Stage**: Runs 6 unit tests
3. **Lint Stage**: Analyzes code quality
4. **Docker Stage**: Creates multi-stage Docker image with reports server
5. **Push Stage**: Publishes image to Docker Hub

### Test Cases

1. Initial counter value is 0
2. Increment increases count by 1
3. Decrement decreases count by 1
4. Multiple increments work correctly
5. Multiple decrements work correctly
6. Mix of increment and decrement operations

### Demonstrating Success and Failure

**Success**: All tests pass (default state)

**Failure**: Modify a test to fail:
```kotlin
// In CounterTest.kt, change:
assertEquals(0, counter.getCount())
// To:
assertEquals(1, counter.getCount())  // Will fail
```

Push the change and watch the CI pipeline fail. Then revert to see it pass.

## Multi-Stage Docker Build

The Dockerfile uses two stages:

**Stage 1 (builder)**: 
- Installs Android SDK and build tools
- Builds the APK
- Runs tests and generates HTML report
- Runs lint and generates HTML report

**Stage 2 (final image)**:
- Lightweight Python base
- Copies only the HTML reports from Stage 1
- Runs Python web server on port 9898
- Final image is much smaller (no Android SDK)

## Technologies Used

- **Language**: Kotlin
- **UI**: Jetpack Compose
- **Build Tool**: Gradle
- **Testing**: JUnit
- **CI/CD**: GitHub Actions
- **Containerization**: Docker (multi-stage build)
- **Report Server**: Python HTTP Server
- **Registry**: Docker Hub

Triggering a build 2