# ��У����ϵͳ API �ӿ��ĵ�

## ? ������Ϣ

- **����URL**: `http://localhost:8080/api/v1`
- **��֤��ʽ**: JWT Bearer Token
- **���ݸ�ʽ**: JSON
- **�ַ�����**: UTF-8
- **���Թ���**: `api_test_utf8.html` (���ò���ҳ��)

## ? ���߲���

**�Ƽ�ʹ�����ò���ҳ��**: �� `api_test_utf8.html` �ļ�������������в�������API�ӿ�

### ����ҳ�湦��
- ? ֧������API�ӿڵ���������
- ? ʵʱ��Ӧ��ʾ�ʹ�����
- ? JWT Token�Զ�����
- ? �ִ������棬������ֱ��
- ? ֧�ֱ������Զ�������֤

## ? ��֤˵��

�󲿷�API�ӿ���Ҫ������ͷ��Я��JWT Token��

```
Authorization: Bearer <your_jwt_token>
```

### ��ȡToken
1. ʹ�ò���ҳ�棺��"��֤�ӿ�"���ֵ��"��¼"��ť
2. ʹ��API��`POST /api/v1/auth/login`
3. ʹ��curl��
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'
```

## ? ͨ����Ӧ��ʽ

### �ɹ���Ӧ
```json
{
    "data": {}, // ��������
    "message": "�����ɹ�",
    "status": "success"
}
```

### ������Ӧ
```json
{
    "error": "������Ϣ",
    "status": "error"
}
```

### ��ҳ��Ӧ
```json
{
    "data": [], // �����б�
    "total": 100, // ������
    "page": 1, // ��ǰҳ��
    "pageSize": 10 // ÿҳ����
}
```

### ���������Ӧ
```json
{
    "status": "ok",
    "message": "Party Building Management System is running"
}
```

---

## 1. ? �û���֤�ӿ�

### 1.1 �û���¼

**�ӿڵ�ַ**: `POST /api/v1/auth/login`

**�������**:
```json
{
    "username": "admin",
    "password": "password"
}
```

**��Ӧʾ��**:
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
        "id": 1,
        "username": "admin",
        "name": "����Ա",
        "role": "admin"
    }
}
```

### 1.2 �û�ע��

**�ӿڵ�ַ**: `POST /api/v1/auth/register`

**�������**:
```json
{
    "username": "newuser",
    "password": "password123",
    "name": "���û�",
    "role": "student" // ��ѡ: admin, teacher, student
}
```

**��Ӧʾ��**:
```json
{
    "message": "User created successfully",
    "user": {
        "id": 4,
        "username": "newuser",
        "name": "���û�",
        "role": "student"
    }
}
```

### 1.3 ��ȡ�û���Ϣ

**�ӿڵ�ַ**: `GET /api/v1/profile`

**����ͷ**:
```
Authorization: Bearer <token>
```

**��Ӧʾ��**:
```json
{
    "user": {
        "id": 1,
        "username": "admin",
        "name": "����Ա",
        "role": "admin"
    }
}
```

---

## 2. ? ѧԱ����ӿ�

### 2.1 ��ȡѧԱ�б�

**�ӿڵ�ַ**: `GET /api/v1/students`

**����ͷ**:
```
Authorization: Bearer <token>
```

**��ѯ����**:
- `page`: ҳ�� (Ĭ��: 1)
- `pageSize`: ÿҳ���� (Ĭ��: 10)

**����ʾ��**:
```
GET /api/v1/students?page=1&pageSize=10
```

**��Ӧʾ��**:
```json
{
    "data": [
        {
            "id": 1,
            "student_no": "2021001",
            "name": "����",
            "gender": "male",
            "ethnicity": "����",
            "birth_date": "2000-01-01T00:00:00Z",
            "education": "����",
            "phone": "13800138001",
            "id_card": "110101200001011234",
            "major_class": "�������ѧ�뼼��",
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

### 2.2 ��ȡѧԱ����

**�ӿڵ�ַ**: `GET /api/v1/students/{id}`

**����ͷ**:
```
Authorization: Bearer <token>
```

**·������**:
- `id`: ѧԱID

**��Ӧʾ��**:
```json
{
    "id": 1,
    "student_no": "2021001",
    "name": "����",
    "gender": "male",
    "ethnicity": "����",
    "birth_date": "2000-01-01T00:00:00Z",
    "education": "����",
    "phone": "13800138001",
    "id_card": "110101200001011234",
    "major_class": "�������ѧ�뼼��",
    "type": "masses",
    "admission_date": "2021-09-01T00:00:00Z",
    "conversion_date": "0001-01-01T00:00:00Z",
    "created_at": "2021-09-01T00:00:00Z",
    "updated_at": "2021-09-01T00:00:00Z"
}
```

### 2.3 ����ѧԱ

**�ӿڵ�ַ**: `POST /api/v1/students`

**����ͷ**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**�������**:
```json
{
    "student_no": "2024001",
    "name": "����",
    "gender": "female",
    "ethnicity": "����",
    "birth_date": "2000-05-15T00:00:00Z",
    "education": "����",
    "phone": "13800138002",
    "id_card": "110101200005151234",
    "major_class": "�������",
    "type": "masses",
    "admission_date": "2024-09-01T00:00:00Z"
}
```

**��Ӧʾ��**:
```json
{
    "id": 4,
    "student_no": "2024001",
    "name": "����",
    "gender": "female",
    "ethnicity": "����",
    "birth_date": "2000-05-15T00:00:00Z",
    "education": "����",
    "phone": "13800138002",
    "id_card": "110101200005151234",
    "major_class": "�������",
    "type": "masses",
    "admission_date": "2024-09-01T00:00:00Z",
    "conversion_date": "0001-01-01T00:00:00Z",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
}
```

### 2.4 ����ѧԱ��Ϣ

**�ӿڵ�ַ**: `PUT /api/v1/students/{id}`

**����ͷ**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**·������**:
- `id`: ѧԱID

**�������**:
```json
{
    "name": "���ģ����£�",
    "gender": "female",
    "type": "activist",
    "phone": "13800138999"
}
```

**��Ӧʾ��**:
```json
{
    "id": 4,
    "student_no": "2024001",
    "name": "���ģ����£�",
    "gender": "female",
    "ethnicity": "����",
    "birth_date": "2000-05-15T00:00:00Z",
    "education": "����",
    "phone": "13800138999",
    "id_card": "110101200005151234",
    "major_class": "�������",
    "type": "activist",
    "admission_date": "2024-09-01T00:00:00Z",
    "conversion_date": "0001-01-01T00:00:00Z",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
}
```

### 2.5 ɾ��ѧԱ

**�ӿڵ�ַ**: `DELETE /api/v1/students/{id}`

**����ͷ**:
```
Authorization: Bearer <token>
```

**·������**:
- `id`: ѧԱID

**��Ӧʾ��**:
```json
{
    "message": "Student deleted successfully"
}
```

---

## 3. ? ��Ա��չ���̽ӿ�

### 3.1 ��ȡѧԱ��չ����

**�ӿڵ�ַ**: `GET /api/v1/students/{id}/development`

**����ͷ**:
```
Authorization: Bearer <token>
```

**·������**:
- `id`: ѧԱID

**��Ӧʾ��**:
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
        "name": "����"
    }
}
```

### 3.2 ����ѧԱ��չ����

**�ӿڵ�ַ**: `PUT /api/v1/students/{id}/development`

**����ͷ**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**·������**:
- `id`: ѧԱID

**�������**:
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

**��Ӧʾ��**:
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

## 4. ? ���ͼ�¼�ӿ�

### 4.1 ��ȡѧԱ���ͼ�¼

**�ӿڵ�ַ**: `GET /api/v1/students/{id}/rewards`

**����ͷ**:
```
Authorization: Bearer <token>
```

**·������**:
- `id`: ѧԱID

**��Ӧʾ��**:
```json
[
    {
        "id": 1,
        "student_id": 1,
        "type": "reward",
        "content": "����ѧ���ɲ�",
        "date": "2021-12-01T00:00:00Z",
        "created_at": "2021-12-01T00:00:00Z",
        "updated_at": "2021-12-01T00:00:00Z",
        "student": {
            "id": 1,
            "student_no": "2021001",
            "name": "����"
        }
    }
]
```

### 4.2 �������ͼ�¼

**�ӿڵ�ַ**: `POST /api/v1/students/{id}/rewards`

**����ͷ**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**·������**:
- `id`: ѧԱID

**�������**:
```json
{
    "type": "reward",
    "content": "����ѧ��",
    "date": "2024-01-01T00:00:00Z"
}
```

**��Ӧʾ��**:
```json
{
    "id": 2,
    "student_id": 1,
    "type": "reward",
    "content": "����ѧ��",
    "date": "2024-01-01T00:00:00Z",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
}
```

---

## 5. ? ��ѵ����ӿ�

### 5.1 ��ȡ��ѵ�б�

**�ӿڵ�ַ**: `GET /api/v1/trainings`

**����ͷ**:
```
Authorization: Bearer <token>
```

**��ѯ����**:
- `page`: ҳ�� (Ĭ��: 1)
- `pageSize`: ÿҳ���� (Ĭ��: 10)

**��Ӧʾ��**:
```json
{
    "data": [
        {
            "id": 1,
            "period": "2021�괺��",
            "student_id": 1,
            "unit": "���˼����ѧԺ",
            "start_date": "2021-03-01T00:00:00Z",
            "end_date": "2021-06-30T00:00:00Z",
            "score": "excellent",
            "certificate_no": "CERT2021001",
            "created_at": "2021-03-01T00:00:00Z",
            "updated_at": "2021-06-30T00:00:00Z",
            "student": {
                "id": 1,
                "student_no": "2021001",
                "name": "����"
            }
        }
    ],
    "total": 1,
    "page": 1
}
```

### 5.2 ��ȡ��ѵ����

**�ӿڵ�ַ**: `GET /api/v1/trainings/{id}`

**����ͷ**:
```
Authorization: Bearer <token>
```

**·������**:
- `id`: ��ѵ��¼ID

**��Ӧʾ��**:
```json
{
    "id": 1,
    "period": "2021�괺��",
    "student_id": 1,
    "unit": "���˼����ѧԺ",
    "start_date": "2021-03-01T00:00:00Z",
    "end_date": "2021-06-30T00:00:00Z",
    "score": "excellent",
    "certificate_no": "CERT2021001",
    "created_at": "2021-03-01T00:00:00Z",
    "updated_at": "2021-06-30T00:00:00Z",
    "student": {
        "id": 1,
        "student_no": "2021001",
        "name": "����"
    }
}
```

### 5.3 ������ѵ��¼

**�ӿڵ�ַ**: `POST /api/v1/trainings`

**����ͷ**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**�������**:
```json
{
    "period": "2024�괺��",
    "student_id": 1,
    "unit": "��У",
    "start_date": "2024-01-01T00:00:00Z",
    "end_date": "2024-01-07T00:00:00Z",
    "score": "excellent",
    "certificate_no": "CERT2024001"
}
```

**��Ӧʾ��**:
```json
{
    "id": 4,
    "period": "2024�괺��",
    "student_id": 1,
    "unit": "��У",
    "start_date": "2024-01-01T00:00:00Z",
    "end_date": "2024-01-07T00:00:00Z",
    "score": "excellent",
    "certificate_no": "CERT2024001",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
}
```

### 5.4 ������ѵ��¼

**�ӿڵ�ַ**: `PUT /api/v1/trainings/{id}`

**����ͷ**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**·������**:
- `id`: ��ѵ��¼ID

**�������**:
```json
{
    "period": "2024�괺�������£�",
    "unit": "��У�����£�",
    "score": "good"
}
```

**��Ӧʾ��**:
```json
{
    "id": 4,
    "period": "2024�괺�������£�",
    "student_id": 1,
    "unit": "��У�����£�",
    "start_date": "2024-01-01T00:00:00Z",
    "end_date": "2024-01-07T00:00:00Z",
    "score": "good",
    "certificate_no": "CERT2024001",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
}
```

### 5.5 ɾ����ѵ��¼

**�ӿڵ�ַ**: `DELETE /api/v1/trainings/{id}`

**����ͷ**:
```
Authorization: Bearer <token>
```

**·������**:
- `id`: ��ѵ��¼ID

**��Ӧʾ��**:
```json
{
    "message": "Training deleted successfully"
}
```

---

## 6. ? ͳ�Ʒ����ӿ�

### 6.1 ��ȡ��ѵͳ��

**�ӿڵ�ַ**: `GET /api/v1/statistics/trainings`

**����ͷ**:
```
Authorization: Bearer <token>
```

**��Ӧʾ��**:
```json
{
    "excellent": 5,
    "good": 8,
    "pass": 12,
    "fail": 2
}
```

---

## 7. ? ϵͳ״̬�ӿ�

### 7.1 �������

**�ӿڵ�ַ**: `GET /health`

**������֤**

**��Ӧʾ��**:
```json
{
    "status": "ok",
    "message": "Party Building Management System is running"
}
```

---

## 8. ? ������˵��

| HTTP״̬�� | ˵�� |
|-----------|------|
| 200 | ����ɹ� |
| 201 | �����ɹ� |
| 400 | ����������� |
| 401 | δ��Ȩ����Ҫ��¼�� |
| 403 | ��ֹ���ʣ�Ȩ�޲��㣩 |
| 404 | ��Դ������ |
| 500 | �������ڲ����� |

## 9. ? �����ֵ�

### 9.1 �û���ɫ (role)
- `admin`: ����Ա
- `teacher`: ��ʦ
- `student`: ѧ��

### 9.2 ѧԱ���� (type)
- `masses`: Ⱥ��
- `activist`: �뵳��������
- `candidate`: ��չ����
- `probationary`: Ԥ����Ա
- `formal`: ��ʽ��Ա

### 9.3 �Ա� (gender)
- `male`: ��
- `female`: Ů

### 9.4 ��ѵ�ɼ� (score)
- `excellent`: ����
- `good`: ����
- `pass`: ����
- `fail`: ������

### 9.5 �������� (type)
- `reward`: ����
- `punishment`: ����

## 10. ? ʹ��ʾ��

### 10.1 ������¼����

```bash
# 1. ��¼��ȡToken
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'

# 2. ʹ��Token�����ܱ����Ľӿ�
curl -X GET http://localhost:8080/api/v1/students \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 10.2 ����ѧԱ��������

```bash
# 1. ��¼��ȡToken
TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}' | jq -r '.token')

# 2. ����ѧԱ
curl -X POST http://localhost:8080/api/v1/students \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "student_no": "2024001",
    "name": "��ѧԱ",
    "gender": "male",
    "type": "masses"
  }'
```

---

## ?? ע������

1. **ʱ���ʽ**: ����ʱ���ֶ�ʹ�� ISO 8601 ��ʽ (RFC3339)
2. **��ҳ**: Ĭ��ÿҳ10����¼�����100��
3. **Token����**: JWT Token��Ч��Ϊ24Сʱ
4. **�ַ�����**: �����ı�ʹ��UTF-8����
5. **������֤**: ��ȷ�������ֶβ�Ϊ��
6. **Ȩ�޿���**: ��ͬ��ɫ�в�ͬ�Ĳ���Ȩ��
7. **���Թ���**: �Ƽ�ʹ�� `api_test_utf8.html` ���нӿڲ���

## ? �������

- **��Ŀ�ֿ�**: [GitHub Repository]
- **API����ҳ��**: `api_test_utf8.html`
- **Postman����**: `Postman_Collection.json`
- **��Ŀ�ĵ�**: `README.md`

---

*������ʱ��: 2025��11��*  
*�ĵ��汾: v2.0*
