# 党校管理系统 API 接口文档

## ? 基础信息

- **基础URL**: `http://localhost:8080/api/v1`
- **认证方式**: JWT Bearer Token
- **数据格式**: JSON
- **字符编码**: UTF-8
- **测试工具**: `api_test_utf8.html` (内置测试页面)

## ? 在线测试

**推荐使用内置测试页面**: 打开 `api_test_utf8.html` 文件即可在浏览器中测试所有API接口

### 测试页面功能
- ? 支持所有API接口的完整测试
- ? 实时响应显示和错误处理
- ? JWT Token自动管理
- ? 现代化界面，操作简单直观
- ? 支持表单数据自动填充和验证

## ? 认证说明

大部分API接口需要在请求头中携带JWT Token：

```
Authorization: Bearer <your_jwt_token>
```

### 获取Token
1. 使用测试页面：在"认证接口"部分点击"登录"按钮
2. 使用API：`POST /api/v1/auth/login`
3. 使用curl：
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'
```

## ? 通用响应格式

### 成功响应
```json
{
    "data": {}, // 具体数据
    "message": "操作成功",
    "status": "success"
}
```

### 错误响应
```json
{
    "error": "错误信息",
    "status": "error"
}
```

### 分页响应
```json
{
    "data": [], // 数据列表
    "total": 100, // 总数量
    "page": 1, // 当前页码
    "pageSize": 10 // 每页数量
}
```

### 健康检查响应
```json
{
    "status": "ok",
    "message": "Party Building Management System is running"
}
```

---

## 1. ? 用户认证接口

### 1.1 用户登录

**接口地址**: `POST /api/v1/auth/login`

**请求参数**:
```json
{
    "username": "admin",
    "password": "password"
}
```

**响应示例**:
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
        "id": 1,
        "username": "admin",
        "name": "管理员",
        "role": "admin"
    }
}
```

### 1.2 用户注册

**接口地址**: `POST /api/v1/auth/register`

**请求参数**:
```json
{
    "username": "newuser",
    "password": "password123",
    "name": "新用户",
    "role": "student" // 可选: admin, teacher, student
}
```

**响应示例**:
```json
{
    "message": "User created successfully",
    "user": {
        "id": 4,
        "username": "newuser",
        "name": "新用户",
        "role": "student"
    }
}
```

### 1.3 获取用户信息

**接口地址**: `GET /api/v1/profile`

**请求头**:
```
Authorization: Bearer <token>
```

**响应示例**:
```json
{
    "user": {
        "id": 1,
        "username": "admin",
        "name": "管理员",
        "role": "admin"
    }
}
```

---

## 2. ? 学员管理接口

### 2.1 获取学员列表

**接口地址**: `GET /api/v1/students`

**请求头**:
```
Authorization: Bearer <token>
```

**查询参数**:
- `page`: 页码 (默认: 1)
- `pageSize`: 每页数量 (默认: 10)

**请求示例**:
```
GET /api/v1/students?page=1&pageSize=10
```

**响应示例**:
```json
{
    "data": [
        {
            "id": 1,
            "student_no": "2021001",
            "name": "张三",
            "gender": "male",
            "ethnicity": "汉族",
            "birth_date": "2000-01-01T00:00:00Z",
            "education": "本科",
            "phone": "13800138001",
            "id_card": "110101200001011234",
            "major_class": "计算机科学与技术",
            "type": "masses",
            "admission_date": "2021-09-01T00:00:00Z",
            "conversion_date": "0001-01-01T00:00:00Z",
            "created_at": "2021-09-01T00:00:00Z",
            "updated_at": "2021-09-01T00:00:00Z"
        }
    ],
    "total": 1,
    "page": 1
}
```

### 2.2 获取学员详情

**接口地址**: `GET /api/v1/students/{id}`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:
- `id`: 学员ID

**响应示例**:
```json
{
    "id": 1,
    "student_no": "2021001",
    "name": "张三",
    "gender": "male",
    "ethnicity": "汉族",
    "birth_date": "2000-01-01T00:00:00Z",
    "education": "本科",
    "phone": "13800138001",
    "id_card": "110101200001011234",
    "major_class": "计算机科学与技术",
    "type": "masses",
    "admission_date": "2021-09-01T00:00:00Z",
    "conversion_date": "0001-01-01T00:00:00Z",
    "created_at": "2021-09-01T00:00:00Z",
    "updated_at": "2021-09-01T00:00:00Z"
}
```

### 2.3 创建学员

**接口地址**: `POST /api/v1/students`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**请求参数**:
```json
{
    "student_no": "2024001",
    "name": "李四",
    "gender": "female",
    "ethnicity": "汉族",
    "birth_date": "2000-05-15T00:00:00Z",
    "education": "本科",
    "phone": "13800138002",
    "id_card": "110101200005151234",
    "major_class": "软件工程",
    "type": "masses",
    "admission_date": "2024-09-01T00:00:00Z"
}
```

**响应示例**:
```json
{
    "id": 4,
    "student_no": "2024001",
    "name": "李四",
    "gender": "female",
    "ethnicity": "汉族",
    "birth_date": "2000-05-15T00:00:00Z",
    "education": "本科",
    "phone": "13800138002",
    "id_card": "110101200005151234",
    "major_class": "软件工程",
    "type": "masses",
    "admission_date": "2024-09-01T00:00:00Z",
    "conversion_date": "0001-01-01T00:00:00Z",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
}
```

### 2.4 更新学员信息

**接口地址**: `PUT /api/v1/students/{id}`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**路径参数**:
- `id`: 学员ID

**请求参数**:
```json
{
    "name": "李四（更新）",
    "gender": "female",
    "type": "activist",
    "phone": "13800138999"
}
```

**响应示例**:
```json
{
    "id": 4,
    "student_no": "2024001",
    "name": "李四（更新）",
    "gender": "female",
    "ethnicity": "汉族",
    "birth_date": "2000-05-15T00:00:00Z",
    "education": "本科",
    "phone": "13800138999",
    "id_card": "110101200005151234",
    "major_class": "软件工程",
    "type": "activist",
    "admission_date": "2024-09-01T00:00:00Z",
    "conversion_date": "0001-01-01T00:00:00Z",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
}
```

### 2.5 删除学员

**接口地址**: `DELETE /api/v1/students/{id}`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:
- `id`: 学员ID

**响应示例**:
```json
{
    "message": "Student deleted successfully"
}
```

---

## 3. ? 党员发展历程接口

### 3.1 获取学员发展历程

**接口地址**: `GET /api/v1/students/{id}/development`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:
- `id`: 学员ID

**响应示例**:
```json
{
    "id": 1,
    "student_id": 1,
    "apply_date": "2021-10-01T00:00:00Z",
    "activist_date": "2021-11-01T00:00:00Z",
    "candidate_date": "0001-01-01T00:00:00Z",
    "probation_date": "0001-01-01T00:00:00Z",
    "conversion_date": "0001-01-01T00:00:00Z",
    "transfer_date": "0001-01-01T00:00:00Z",
    "introduction_date": "0001-01-01T00:00:00Z",
    "status": "masses",
    "created_at": "2021-10-01T00:00:00Z",
    "updated_at": "2021-11-01T00:00:00Z",
    "student": {
        "id": 1,
        "student_no": "2021001",
        "name": "张三"
    }
}
```

### 3.2 更新学员发展历程

**接口地址**: `PUT /api/v1/students/{id}/development`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**路径参数**:
- `id`: 学员ID

**请求参数**:
```json
{
    "apply_date": "2021-10-01T00:00:00Z",
    "activist_date": "2021-11-01T00:00:00Z",
    "candidate_date": "2022-03-01T00:00:00Z",
    "probation_date": "2022-06-01T00:00:00Z",
    "conversion_date": "2023-06-01T00:00:00Z",
    "status": "formal"
}
```

**响应示例**:
```json
{
    "id": 1,
    "student_id": 1,
    "apply_date": "2021-10-01T00:00:00Z",
    "activist_date": "2021-11-01T00:00:00Z",
    "candidate_date": "2022-03-01T00:00:00Z",
    "probation_date": "2022-06-01T00:00:00Z",
    "conversion_date": "2023-06-01T00:00:00Z",
    "transfer_date": "0001-01-01T00:00:00Z",
    "introduction_date": "0001-01-01T00:00:00Z",
    "status": "formal",
    "created_at": "2021-10-01T00:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
}
```

---

## 4. ? 奖惩记录接口

### 4.1 获取学员奖惩记录

**接口地址**: `GET /api/v1/students/{id}/rewards`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:
- `id`: 学员ID

**响应示例**:
```json
[
    {
        "id": 1,
        "student_id": 1,
        "type": "reward",
        "content": "优秀学生干部",
        "date": "2021-12-01T00:00:00Z",
        "created_at": "2021-12-01T00:00:00Z",
        "updated_at": "2021-12-01T00:00:00Z",
        "student": {
            "id": 1,
            "student_no": "2021001",
            "name": "张三"
        }
    }
]
```

### 4.2 创建奖惩记录

**接口地址**: `POST /api/v1/students/{id}/rewards`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**路径参数**:
- `id`: 学员ID

**请求参数**:
```json
{
    "type": "reward",
    "content": "三好学生",
    "date": "2024-01-01T00:00:00Z"
}
```

**响应示例**:
```json
{
    "id": 2,
    "student_id": 1,
    "type": "reward",
    "content": "三好学生",
    "date": "2024-01-01T00:00:00Z",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
}
```

---

## 5. ? 培训管理接口

### 5.1 获取培训列表

**接口地址**: `GET /api/v1/trainings`

**请求头**:
```
Authorization: Bearer <token>
```

**查询参数**:
- `page`: 页码 (默认: 1)
- `pageSize`: 每页数量 (默认: 10)

**响应示例**:
```json
{
    "data": [
        {
            "id": 1,
            "period": "2021年春季",
            "student_id": 1,
            "unit": "马克思主义学院",
            "start_date": "2021-03-01T00:00:00Z",
            "end_date": "2021-06-30T00:00:00Z",
            "score": "excellent",
            "certificate_no": "CERT2021001",
            "created_at": "2021-03-01T00:00:00Z",
            "updated_at": "2021-06-30T00:00:00Z",
            "student": {
                "id": 1,
                "student_no": "2021001",
                "name": "张三"
            }
        }
    ],
    "total": 1,
    "page": 1
}
```

### 5.2 获取培训详情

**接口地址**: `GET /api/v1/trainings/{id}`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:
- `id`: 培训记录ID

**响应示例**:
```json
{
    "id": 1,
    "period": "2021年春季",
    "student_id": 1,
    "unit": "马克思主义学院",
    "start_date": "2021-03-01T00:00:00Z",
    "end_date": "2021-06-30T00:00:00Z",
    "score": "excellent",
    "certificate_no": "CERT2021001",
    "created_at": "2021-03-01T00:00:00Z",
    "updated_at": "2021-06-30T00:00:00Z",
    "student": {
        "id": 1,
        "student_no": "2021001",
        "name": "张三"
    }
}
```

### 5.3 创建培训记录

**接口地址**: `POST /api/v1/trainings`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**请求参数**:
```json
{
    "period": "2024年春季",
    "student_id": 1,
    "unit": "党校",
    "start_date": "2024-01-01T00:00:00Z",
    "end_date": "2024-01-07T00:00:00Z",
    "score": "excellent",
    "certificate_no": "CERT2024001"
}
```

**响应示例**:
```json
{
    "id": 4,
    "period": "2024年春季",
    "student_id": 1,
    "unit": "党校",
    "start_date": "2024-01-01T00:00:00Z",
    "end_date": "2024-01-07T00:00:00Z",
    "score": "excellent",
    "certificate_no": "CERT2024001",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
}
```

### 5.4 更新培训记录

**接口地址**: `PUT /api/v1/trainings/{id}`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**路径参数**:
- `id`: 培训记录ID

**请求参数**:
```json
{
    "period": "2024年春季（更新）",
    "unit": "党校（更新）",
    "score": "good"
}
```

**响应示例**:
```json
{
    "id": 4,
    "period": "2024年春季（更新）",
    "student_id": 1,
    "unit": "党校（更新）",
    "start_date": "2024-01-01T00:00:00Z",
    "end_date": "2024-01-07T00:00:00Z",
    "score": "good",
    "certificate_no": "CERT2024001",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
}
```

### 5.5 删除培训记录

**接口地址**: `DELETE /api/v1/trainings/{id}`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:
- `id`: 培训记录ID

**响应示例**:
```json
{
    "message": "Training deleted successfully"
}
```

---

## 6. ? 统计分析接口

### 6.1 获取培训统计

**接口地址**: `GET /api/v1/statistics/trainings`

**请求头**:
```
Authorization: Bearer <token>
```

**响应示例**:
```json
{
    "excellent": 5,
    "good": 8,
    "pass": 12,
    "fail": 2
}
```

---

## 7. ? 系统状态接口

### 7.1 健康检查

**接口地址**: `GET /health`

**无需认证**

**响应示例**:
```json
{
    "status": "ok",
    "message": "Party Building Management System is running"
}
```

---

## 8. ? 错误码说明

| HTTP状态码 | 说明 |
|-----------|------|
| 200 | 请求成功 |
| 201 | 创建成功 |
| 400 | 请求参数错误 |
| 401 | 未授权（需要登录） |
| 403 | 禁止访问（权限不足） |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

## 9. ? 数据字典

### 9.1 用户角色 (role)
- `admin`: 管理员
- `teacher`: 教师
- `student`: 学生

### 9.2 学员类型 (type)
- `masses`: 群众
- `activist`: 入党积极分子
- `candidate`: 发展对象
- `probationary`: 预备党员
- `formal`: 正式党员

### 9.3 性别 (gender)
- `male`: 男
- `female`: 女

### 9.4 培训成绩 (score)
- `excellent`: 优秀
- `good`: 良好
- `pass`: 及格
- `fail`: 不及格

### 9.5 奖惩类型 (type)
- `reward`: 奖励
- `punishment`: 处分

## 10. ? 使用示例

### 10.1 完整登录流程

```bash
# 1. 登录获取Token
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'

# 2. 使用Token访问受保护的接口
curl -X GET http://localhost:8080/api/v1/students \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 10.2 创建学员完整流程

```bash
# 1. 登录获取Token
TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}' | jq -r '.token')

# 2. 创建学员
curl -X POST http://localhost:8080/api/v1/students \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "student_no": "2024001",
    "name": "新学员",
    "gender": "male",
    "type": "masses"
  }'
```

---

## ?? 注意事项

1. **时间格式**: 所有时间字段使用 ISO 8601 格式 (RFC3339)
2. **分页**: 默认每页10条记录，最大100条
3. **Token过期**: JWT Token有效期为24小时
4. **字符编码**: 所有文本使用UTF-8编码
5. **数据验证**: 请确保必填字段不为空
6. **权限控制**: 不同角色有不同的操作权限
7. **测试工具**: 推荐使用 `api_test_utf8.html` 进行接口测试

## ? 相关链接

- **项目仓库**: [GitHub Repository]
- **API测试页面**: `api_test_utf8.html`
- **Postman集合**: `Postman_Collection.json`
- **项目文档**: `README.md`

---

*最后更新时间: 2025年11月*  
*文档版本: v2.0*
