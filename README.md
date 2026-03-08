[Tiếng Việt](README.md) | [English](README-eng.md)

# ☸️ Karma Monitor (Hộ Pháp Auditor)
![Go Version](https://img.shields.io/badge/Go-1.25.4-blue?logo=go)
![Docker](https://img.shields.io/badge/Docker-Enabled-brightgreen?logo=docker)
![Security](https://img.shields.io/badge/Security-Full%20Giáp-red)
![Karma](https://img.shields.io/badge/Karma-Audit%20Only-orange)
---
Karma Monitor là một công cụ giám sát "nghiệp lực" thời gian thực trên không gian số. Hệ thống được xây dựng bằng Go (Golang) để đảm bảo hiệu năng "Kim Cương bất hoại" và khả năng truy vết không bỏ sót một nhịp push nào.
---
## ✨ Tính năng "Hộ Pháp"
- **Real-time Polling:** Tự động rình rập GitHub Public Events với độ trễ cực thấp.
- **Deep Secret Scan:** Truy quét "nghiệp" (secrets, env, passwords) ẩn nấp trong các commit diff bằng Regex chuẩn Security.
- **Spiritual Persistence:** Cơ chế lưu trữ ID cuối cùng vào "Sổ bìa đen" (state file) để đảm bảo dù có ngủ gật (System Sleep) vẫn không bỏ lỡ drama.
- **Ò Í E Alert:** Thông báo tức thời qua Discord Webhook mỗi khi phát hiện "hề kịch câm" diễn ra.
---
## 🛠️ Công nghệ "Vạn Năng"
- **Language:** Go 1.25.4 (Tối ưu cho M4 Pro).
- **Container:** Docker & Docker Compose (Cách ly hoàn toàn với bụi trần).
- **Infrastructure:** Multi-stage Build (Nhẹ như lông hồng, chỉ ~15MB).
---
## 🚀 Hướng dẫn "Thỉnh" Bot về máy
- Clone repo này về con máy triệu đô của thí chủ.
- Chỉnh sửa main.go (Điền username và Webhook vào "vùng kín").
- Khởi tạo sổ bìa đen: `mkdir state`.
- Triển khai: `docker compose up -d --build`.
---
# Lưu ý: Công cụ này dùng để "Học làm người", không khuyến khích các hành vi hãm tài. 🙏
