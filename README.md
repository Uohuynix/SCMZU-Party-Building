# ��У����ϵͳ

**����ʱ��**: 2025��8�¡���2025��11��  
**�����Ŷ�**: ��˼·ʵ����Geek��2023��ȫ��

����Ŀ����Go���Ժ�Gin��ܿ����ĵ�У����ϵͳ���ṩѧԱ��Ϣ������ѵ��¼������Ա��չ���̹���ȹ��ܡ�ϵͳ����RESTful API��ƣ�֧��������CRUD������JWT��֤��

## ? ��������

**���߲���**: �� `api_test_utf8.html` �ļ�������������в�������API�ӿ�
- ֧��������֤��ѧԱ������ѵ����Ƚӿ�
- ʵʱ��Ӧ��ʾ�ʹ�����
- �ִ������棬������ֱ��

## ? ��Ҫ����

- ? **�û���֤����Ȩ** - JWT Token��֤��֧�ֶ��ɫȨ�޿���
- ? **ѧԱ��Ϣ����** - ������ѧԱCRUD������֧�ַ�ҳ��ѯ
- ? **��ѵ��¼����** - ��ѵ�γ̼�¼���ɼ������֤�����
- ? **��Ա��չ����** - ��Ⱥ�ڵ���ʽ��Ա��������չ�켣
- ? **���ͼ�¼����** - ѧԱ�����ʹ��ּ�¼����
- ? **����ͳ�Ʒ���** - ��ѵ�ɼ�ͳ�ƺ����ݷ���
- ? **RESTful API** - ��׼����API�ӿ����
- ? **���߲��Թ���** - ����API����ҳ�棬֧��ʵʱ����

## ?? ����ջ

- **��˿��**: Gin (Go Web���)
- **���ݿ�**: MySQL 5.7+
- **ORM**: GORM (Go ORM��)
- **��֤**: JWT (JSON Web Token)
- **�������**: Go 1.21+
- **ǰ�˲���**: HTML5 + JavaScript (API����ҳ��)

## ? ��Ŀ�ṹ

```
dangjiantest/
������ api/                    # API��������
��   ������ auth_handler.go     # ��֤��ؽӿ�
��   ������ student_handler.go  # ѧԱ����ӿ�
��   ������ training_handler.go # ��ѵ����ӿ�
��   ������ stats_handler.go    # ͳ�Ʒ����ӿ�
������ config/                 # ���ù���
��   ������ config.go          # ���ýṹ����
������ dao/                   # ���ݷ��ʲ�
��   ������ base_dao.go        # ����DAO
��   ������ user_dao.go        # �û����ݷ���
��   ������ student_dao.go     # ѧԱ���ݷ���
��   ������ training_dao.go    # ��ѵ���ݷ���
��   ������ develop_dao.go     # ��չ�������ݷ���
��   ������ reward_dao.go      # ���ͼ�¼���ݷ���
������ entity/                # ����ģ��
��   ������ user.go            # �û�ʵ��
��   ������ student.go         # ѧԱʵ��
��   ������ training.go        # ��ѵʵ��
��   ������ development.go     # ��չ����ʵ��
��   ������ reward.go          # ���ͼ�¼ʵ��
������ middleware/            # �м��
��   ������ auth.go            # JWT��֤�м��
������ pkg/                   # ������
��   ������ jwt/
��       ������ jwt.go         # JWT���ߺ���
������ service/               # ҵ���߼���
��   ������ auth_service.go    # ��֤ҵ���߼�
��   ������ student_service.go # ѧԱ����ҵ���߼�
��   ������ training_service.go# ��ѵ����ҵ���߼�
��   ������ stats_service.go   # ͳ�Ʒ���ҵ���߼�
������ util/                  # ���ߺ���
��   ������ password.go        # ������ܹ���
������ db/                    # ���ݿ����
��   ������ db.go             # ���ݿ�����
��   ������ restore.go        # ���ݻָ�
������ main.go               # �������
������ api_test_utf8.html    # API����ҳ�� ?
������ API_DOCUMENTATION.md  # API�ӿ��ĵ�
������ config.env            # ��������ʾ��
������ start_server.bat      # Windows�����ű�
������ go.mod                # Goģ���ļ�
������ README.md             # ��Ŀ˵���ĵ�
```

## ? ����Ҫ��

- **Go**: 1.21 ����߰汾
- **MySQL**: 5.7 ����߰汾
- **�����**: ֧��HTML5��JavaScript���ִ������������API���ԣ�

## ? ���ٿ�ʼ

### 1. ��¡��Ŀ

```bash
git clone <repository-url>
cd dangjiantest
```

### 2. ��װ����

```bash
go mod tidy
```

### 3. ���û�������

���� `config.env` �ļ����޸����ݿ����ã�

```bash
# ���ݿ�����
DB_HOST=localhost
DB_PORT=3306
DB_USER=����û���
DB_PASSWORD=�������
DB_NAME=party_school

# JWT��Կ
JWT_SECRET=your_jwt_secret_key

# ����������
APP_PORT=8080
LOG_LEVEL=info
ENVIRONMENT=development
```

### 4. �������ݿ�

```sql
CREATE DATABASE party_school CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 5. ����������

#### ����һ��ʹ�������ű����Ƽ���

```bash
# Windows
start_server.bat

# Linux/Mac
chmod +x start_server.sh
./start_server.sh
```

#### ���������ֶ�����

```bash
go run main.go
```

#### �����������������

```bash
go build -o party_school.exe .
./party_school.exe
```

### 6. ����API

������������

1. **�������**: ���� `http://localhost:8080/health`
2. **API����**: �� `api_test_utf8.html` �ļ�������������
3. **API�ĵ�**: �鿴 `API_DOCUMENTATION.md` �˽���ϸ�ӿ�

### 7. Ĭ���˺�

- **����Ա**: `admin` / `password`
- **��ʦ**: `teacher` / `password`
- **ѧ��**: `student` / `password`

## ? API�ӿ��ĵ�

### ? ���߲���

**�Ƽ�ʹ�����ò���ҳ��**: �� `api_test_utf8.html` �ļ�������������в�������API�ӿ�

### ���Ľӿڸ���

#### ? �û���֤
- `POST /api/v1/auth/login` - �û���¼
- `POST /api/v1/auth/register` - �û�ע��  
- `GET /api/v1/profile` - ��ȡ�û���Ϣ

#### ? ѧԱ����
- `GET /api/v1/students` - ��ȡѧԱ�б���ҳ��
- `GET /api/v1/students/{id}` - ��ȡѧԱ����
- `POST /api/v1/students` - ����ѧԱ
- `PUT /api/v1/students/{id}` - ����ѧԱ��Ϣ
- `DELETE /api/v1/students/{id}` - ɾ��ѧԱ

#### ? ��չ����
- `GET /api/v1/students/{id}/development` - ��ȡ��չ����
- `PUT /api/v1/students/{id}/development` - ���·�չ����

#### ? ���ͼ�¼
- `GET /api/v1/students/{id}/rewards` - ��ȡ���ͼ�¼
- `POST /api/v1/students/{id}/rewards` - �������ͼ�¼

#### ? ��ѵ����
- `GET /api/v1/trainings` - ��ȡ��ѵ�б���ҳ��
- `GET /api/v1/trainings/{id}` - ��ȡ��ѵ����
- `POST /api/v1/trainings` - ������ѵ��¼
- `PUT /api/v1/trainings/{id}` - ������ѵ��¼
- `DELETE /api/v1/trainings/{id}` - ɾ����ѵ��¼

#### ? ͳ�Ʒ���
- `GET /api/v1/statistics/trainings` - ��ȡ��ѵͳ��

#### ? ϵͳ״̬
- `GET /health` - �������

### ��ϸ�ĵ�

������API�ӿ��ĵ���鿴��[API_DOCUMENTATION.md](API_DOCUMENTATION.md)

## ���ݿ�ṹ

### �û��� (users)
- id: ����
- username: �û�����Ψһ��
- password: ���루���ܣ�
- role: ��ɫ��admin/teacher/student��
- name: ����
- created_at: ����ʱ��
- updated_at: ����ʱ��

### ѧԱ�� (students)
- id: ����
- student_no: ѧ�ţ�Ψһ��
- name: ����
- gender: �Ա���/Ů��
- ethnicity: ����
- birth_date: ��������
- education: ѧ��
- phone: �绰
- id_card: ���֤��
- major_class: רҵ�༶
- type: ������ò��Ⱥ��/�뵳��������/��չ����/Ԥ����Ա/��ʽ��Ա��
- admission_date: ��ѧ����
- conversion_date: ת������
- created_at: ����ʱ��
- updated_at: ����ʱ��

### ��ѵ�� (trainings)
- id: ����
- period: �ڴ�
- student_id: ѧԱID�������
- unit: ��ѵ��λ
- start_date: ��ʼ����
- end_date: ��������
- score: �ɼ�������/����/����/������
- certificate_no: ֤����
- created_at: ����ʱ��
- updated_at: ����ʱ��

### ��չ���̱� (developments)
- id: ����
- student_id: ѧԱID�������
- apply_date: �����뵳����
- activist_date: �뵳������������
- candidate_date: ��չ��������
- probation_date: Ԥ����Ա����
- conversion_date: ת������
- transfer_date: ת������
- introduction_date: ����������
- status: ��ǰ״̬��Ⱥ��/�뵳��������/��չ����/Ԥ����Ա/��ʽ��Ա��
- created_at: ����ʱ��
- updated_at: ����ʱ��

### ���ͱ� (rewards)
- id: ����
- student_id: ѧԱID�������
- type: ���ͣ�����/���֣�
- content: ����
- date: ����
- created_at: ����ʱ��
- updated_at: ����ʱ��

## ����ָ��

### ����¹���

1. �� `entity/` �ж�������ģ��
2. �� `dao/` ��ʵ�����ݷ����߼�
3. �� `service/` ��ʵ��ҵ���߼�
4. �� `api/` ��ʵ��HTTP�ӿ�
5. �� `main.go` ��ע��·��

### ����淶

- ��ѭGo���Ա���淶
- ʹ�� `go fmt` ��ʽ������
- ��ӱ�Ҫ��ע��
- ��д��Ԫ����

## �����ų�

### ��������

1. **Go����������**
   - ������� "package xxx is not in std" ��������Go��װ
   - �������°�װGo��ʹ��ϵͳPATH�е�Go

2. **���ݿ���������**
   - ���MySQL�����Ƿ�����
   - ��֤���ݿ�������Ϣ
   - ȷ�����ݿ��û�Ȩ��

3. **�˿�ռ������**
   - �޸� `config.env` �е� `SERVER_PORT` ����
   - ��ʹ�� `netstat -ano | findstr :8080` ����ռ�ý���

### ��־�鿴

��������ʱ�������ϸ����־��Ϣ��������
- ���ݿ�����״̬
- ������������Ϣ
- API������־
- ������Ϣ

## ? ����ĵ�

- **[API�ӿ��ĵ�](API_DOCUMENTATION.md)** - ������API�ӿ�˵����ʹ��ʾ��
- **[API����ҳ��](api_test_utf8.html)** - ���õ����߲��Թ��� ?
- **[Postman����](Postman_Collection.json)** - �ɵ���Postman�Ľӿڼ���
- **[Apifox����](apifox_collection.json)** - �ɵ���Apifox�Ľӿڼ���
- **[OpenAPI 3.0](api_openapi.yaml)** - ��׼OpenAPI 3.0��ʽ�ĵ�
- **[Apifox�����ļ�](apifox_import.json)** - רΪApifox�Ż��ĵ����ļ�
- **[Protocol Buffers](api.proto)** - gRPC/Protocol Buffers�����ļ�

## ?? ����ָ��

### ����¹���

1. �� `entity/` �ж�������ģ��
2. �� `dao/` ��ʵ�����ݷ����߼�
3. �� `service/` ��ʵ��ҵ���߼�
4. �� `api/` ��ʵ��HTTP�ӿ�
5. �� `main.go` ��ע��·��
6. ���� `api_test_utf8.html` ����ҳ��

### ����淶

- ��ѭGo���Ա���淶
- ʹ�� `go fmt` ��ʽ������
- ��ӱ�Ҫ��ע��
- ��д��Ԫ����

## ? �����ų�

### ��������

1. **Go����������**
   - ���Go��װ�Ͱ汾
   - �������°�װGo��ʹ��ϵͳPATH�е�Go

2. **���ݿ���������**
   - ���MySQL�����Ƿ�����
   - ��֤���ݿ�������Ϣ
   - ȷ�����ݿ��û�Ȩ��

3. **�˿�ռ������**
   - �޸� `config.env` �е� `APP_PORT` ����
   - ʹ�� `netstat -ano | findstr :8080` ����ռ�ý���

4. **����ҳ����������**
   - ȷ��ʹ�� `api_test_utf8.html` �ļ�
   - ����������������

## ? ���֤

MIT License

## ? ����

��ӭ�ύIssue��Pull Request���Ľ���Ŀ

## ? ��ϵ��ʽ

- **�����Ŷ�**: ��˼·ʵ����Geek��2023��ȫ��
- **��Ŀ��ַ**: [GitHub Repository]
- **���ⷴ��**: ��ͨ��GitHub Issues�ύ
