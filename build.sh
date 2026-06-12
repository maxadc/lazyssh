#!/bin/bash
set -e

BINARY_NAME="lazyssh"
OUTPUT_DIR="./bin"
CMD_DIR="./cmd"
VERSION=${VERSION:-"v0.1.0"}
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME=$(date +"%Y-%m-%d %H:%M:%S")
BUILD_TAGS=""

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

info() { echo -e "${GREEN}[INFO]${NC} $1"; }
warn() { echo -e "${YELLOW}[WARN]${NC} $1"; }
error() { echo -e "${RED}[ERROR]${NC} $1"; exit 1; }

check_go() {
    if ! command -v go &> /dev/null; then
        error "Go is not installed. Please install Go first."
    fi
    info "Go version: $(go version)"
}

download_deps() {
    info "Downloading dependencies..."
    go mod download
    go mod verify
    info "Dependencies downloaded successfully."
}

run_tests() {
    info "Running tests..."
    if go test -race -short ./...; then
        info "All tests passed."
    else
        warn "Some tests failed, but continuing with build..."
    fi
}

build_binary() {
    local os=$1
    local arch=$2
    local output=$3
    
    info "Building for $os/$arch..."
    CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build \
        -ldflags "-X main.version=$VERSION -X main.gitCommit=$GIT_COMMIT -s -w" \
        -tags "$BUILD_TAGS" \
        -o "$OUTPUT_DIR/$output" \
        "$CMD_DIR"
    info "Built: $OUTPUT_DIR/$output"
}

build_current() {
    info "Building for current platform..."
    go build \
        -ldflags "-X main.version=$VERSION -X main.gitCommit=$GIT_COMMIT" \
        -tags "$BUILD_TAGS" \
        -o "$OUTPUT_DIR/$BINARY_NAME" \
        "$CMD_DIR"
    info "Built: $OUTPUT_DIR/$BINARY_NAME"
}

build_all() {
    info "Building for all platforms..."
    build_binary "linux" "amd64" "$BINARY_NAME-linux-amd64"
    build_binary "linux" "arm64" "$BINARY_NAME-linux-arm64"
    build_binary "darwin" "amd64" "$BINARY_NAME-darwin-amd64"
    build_binary "darwin" "arm64" "$BINARY_NAME-darwin-arm64"
    build_binary "windows" "amd64" "$BINARY_NAME-windows-amd64.exe"
}

clean() {
    info "Cleaning build artifacts..."
    rm -rf "$OUTPUT_DIR"
    go clean -cache -testcache
    info "Clean completed."
}

show_help() {
    echo "Usage: $0 [command]"
    echo ""
    echo "Commands:"
    echo "  build       Build for current platform (English, default)"
    echo "  zh          Build for current platform (Chinese)"
    echo "  all         Build for all platforms"
    echo "  all-zh      Build for all platforms (Chinese)"
    echo "  test        Run tests only"
    echo "  clean       Clean build artifacts"
    echo "  help        Show this help message"
    echo ""
    echo "Environment Variables:"
    echo "  VERSION     Set version (default: v0.1.0)"
}

mkdir -p "$OUTPUT_DIR"

case "${1:-build}" in
    build)
        check_go
        download_deps
        build_current
        info "Build completed successfully!"
        ;;
    zh)
        BUILD_TAGS="zh"
        check_go
        download_deps
        build_current
        info "Chinese version build completed successfully!"
        ;;
    build-zh)
        BUILD_TAGS="zh"
        check_go
        download_deps
        build_current
        info "Chinese version build completed successfully!"
        ;;
    all)
        check_go
        download_deps
        build_all
        info "All builds completed successfully!"
        ;;
    all-zh)
        BUILD_TAGS="zh"
        check_go
        download_deps
        build_all
        info "Chinese version all builds completed successfully!"
        ;;
    test)
        check_go
        run_tests
        ;;
    clean)
        clean
        ;;
    help|--help|-h)
        show_help
        ;;
    *)
        error "Unknown command: $1. Use 'help' for usage."
        ;;
esac