# Stage 1: Build Environment
FROM openjdk:26-ea-17-jdk-slim AS builder

ENV ANDROID_SDK_ROOT=/opt/android-sdk
ENV PATH=${PATH}:${ANDROID_SDK_ROOT}/cmdline-tools/latest/bin:${ANDROID_SDK_ROOT}/platform-tools

# Install dependencies
RUN apt-get update && apt-get install -y \
    wget \
    unzip \
    && rm -rf /var/lib/apt/lists/*

# Install Android SDK
RUN mkdir -p ${ANDROID_SDK_ROOT}/cmdline-tools && \
    wget -q https://dl.google.com/android/repository/commandlinetools-linux-9477386_latest.zip -O /tmp/cmdline-tools.zip && \
    unzip -q /tmp/cmdline-tools.zip -d ${ANDROID_SDK_ROOT}/cmdline-tools && \
    mv ${ANDROID_SDK_ROOT}/cmdline-tools/cmdline-tools ${ANDROID_SDK_ROOT}/cmdline-tools/latest && \
    rm /tmp/cmdline-tools.zip

# Accept licenses and install required SDK components
RUN yes | sdkmanager --licenses && \
    sdkmanager "platform-tools" "platforms;android-31" "build-tools;31.0.0"

WORKDIR /app

# Copy project files
COPY . .

# Grant execute permission
RUN chmod +x ./gradlew

# Build APK
RUN ./gradlew assembleDebug --no-daemon

# Run tests and generate test reports
RUN ./gradlew test --no-daemon

# Run lint and generate lint reports
RUN ./gradlew lint --no-daemon

# Stage 2: Clone Server Repository
FROM alpine/git AS server-repo

# Clone the report server repository
# Replace 'yourusername' with your actual GitHub username
RUN git clone https://github.com/NikhilShankar/AndroidReports.git /server

# Stage 3: Report Server
FROM python:3.11-slim

WORKDIR /reports

# Copy test reports from builder stage
COPY --from=builder /app/app/build/reports/tests/testDebugUnitTest/ /reports/testresults/

# Copy lint reports from builder stage
COPY --from=builder /app/app/build/reports/lint-results-debug.html /reports/lint/index.html

# Copy server script from cloned repository
COPY --from=server-repo /server/server.py /reports/server.py

EXPOSE 9898

CMD ["python", "server.py"]