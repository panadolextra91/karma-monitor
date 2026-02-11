# Stage 1: Build - Dùng golang alpine cho nhẹ
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
# Build ra binary "ho-phap"
RUN go build -o ho-phap main.go

# Stage 2: Run - Chỉ lấy cái binary thôi, vứt hết rác build đi
FROM alpine:latest
WORKDIR /root/
# Copy binary từ stage builder qua
COPY --from=builder /app/ho-phap .
# Cài thêm ca-certificates để gọi được HTTPS (API GitHub/Discord)
RUN apk --no-cache add ca-certificates

# Khởi động Hộ Pháp
CMD ["./ho-phap"]