# Go Fiber Todo API

โปรเจกต์ตัวอย่าง Todo API backend ด้วย Go, Fiber และ GORM พร้อมเชื่อมต่อ PostgreSQL

---

## สเปคระบบ

- **Go version**: 1.23.0 (Toolchain: go1.24.5)
- **Module**: `go-fiber-todo`
- **Web Framework**: Fiber v2 (เร็ว เบา คล้าย Express.js)
- **Database**: PostgreSQL + GORM ORM
- **Testing**: Testify framework

---

## Dependencies

| Package                       | Version | คำอธิบาย                                                  |
| ----------------------------- | ------- | --------------------------------------------------------- |
| `github.com/gofiber/fiber/v2` | v2.52.9 | Web framework เบา เร็ว คล้าย Express.js สำหรับ Go         |
| `gorm.io/gorm`                | v1.30.0 | ORM (Object Relational Mapper) สำหรับจัดการฐานข้อมูลใน Go |
| `gorm.io/driver/postgres`     | v1.6.0  | PostgreSQL driver สำหรับ GORM                             |
| `github.com/stretchr/testify` | v1.8.4  | ไลบรารีช่วยเขียน unit test และ assertion ใน Go            |

---

## การติดตั้ง

### 1. ข้อกำหนดเบื้องต้น

- Go เวอร์ชัน 1.23.0 หรือสูงกว่า
- PostgreSQL Database (Local หรือ Cloud เช่น Neon)

### 2. ติดตั้งโปรเจกต์

```bash
git clone <repository-url>
cd go-fiber-todo
go mod tidy
```

### 3. ตั้งค่าฐานข้อมูล

แก้ไขตัวแปร DSN ในโค้ดหรือไฟล์ `.env`:

```
DATABASE_URL="postgresql://username:password@host:port/dbname?sslmode=require"
```

### 4. รันโปรเจกต์

```bash
# รันแบบปกติ
go run main.go

# รันด้วย hot reload (ต้องติดตั้ง air ก่อน)
air
```

Server จะรันที่: `http://localhost:3000`

---

## โครงสร้างโปรเจกต์

```
go-fiber-todo/
├── main.go                 # Entry point ของแอปพลิเคชัน
├── user/                   # User module
│   ├── user.controller.go  # User controllers (handlers)
│   ├── user.service.go     # User business logic
│   └── user.model.go       # User data models
├── todo/                   # Todo module
│   ├── todo.controller.go  # Todo controllers (handlers)
│   ├── todo.service.go     # Todo business logic
│   └── todo.model.go       # Todo data models
├── config/                 # Configuration files
│   └── database.go         # Database configuration
├── middleware/             # Custom middlewares
│   └── cors.go             # CORS middleware
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies checksum
├── .env.example            # Environment variables template
└── .air.toml               # Air configuration (hot reload)
```

---

## API Endpoints

### Users

- `GET /api/v1/users` - ดึงรายการ users ทั้งหมด
- `POST /api/v1/users` - สร้าง user ใหม่
- `GET /api/v1/users/:id` - ดึงข้อมูล user ตาม ID
- `PUT /api/v1/users/:id` - อัปเดตข้อมูล user
- `DELETE /api/v1/users/:id` - ลบ user

### Todos

- `GET /api/v1/todos` - ดึงรายการ todos ทั้งหมด
- `POST /api/v1/todos` - สร้าง todo ใหม่
- `GET /api/v1/todos/:id` - ดึงข้อมูล todo ตาม ID
- `PUT /api/v1/todos/:id` - อัปเดตข้อมูล todo
- `DELETE /api/v1/todos/:id` - ลบ todo

---

## การทดสอบ

### รันเทสต์

```bash
go test ./...
```

### รันเทสต์พร้อม coverage

```bash
go test -cover ./...
```

---

## Features

- ✅ RESTful API design
- ✅ CRUD operations สำหรับ Users และ Todos
- ✅ Database migration อัตโนมัติ
- ✅ Error handling และ validation
- ✅ CORS support
- ✅ Request logging
- ✅ Hot reload support (Air)
- ✅ Unit testing support

---

## การใช้งาน

1. **เริ่มต้น**: สร้าง user ใหม่ผ่าน POST `/api/v1/users`
2. **จัดการ Todos**: สร้าง, อ่าน, อัปเดต, ลบ todos ผ่าน API endpoints
3. **ความสัมพันธ์**: Todo แต่ละรายการจะเชื่อมโยงกับ User ผ่าน `user_id`

---

## สรุป

- ใช้ Go + Fiber ทำ API server
- ใช้ GORM เป็น ORM เชื่อมต่อ PostgreSQL
- มี unit testing support ด้วย testify
- ใช้ air ช่วย hot reload ระหว่างพัฒนา
- โครงสร้างโค้ดแบบ modular และเป็นระเบียบ
