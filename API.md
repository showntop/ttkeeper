# ttkeeper application interface document

- devp host: localhost:8089
- test host: localhost:8089
- prod host: localhost:8089

### 1.sess 登录/注销
	post /ss
	{
		"username": string,
		"password": string,
	}
	delete /ss

### 1.mime
	get /me/p

### 1.users 用户
	post /u
	{
		"username": string,
		"password": string,
		"orgunit_id": number
	}
	例：
	{
		"username": "u2",
		"password": "u2234567890",
		"orgunit_id": 2,
	}
	
	get /u

### 1.roles 角色
	post /r
	{
		"name": string,
	}
	例：
	{
		"name": "管理员",
	}
	
	get /r

### 1.grant roles to user
	post /ur
	{
		"user_id": number,
		"role_id": number,
	}
	
	get /ur

### 1.resources 资源
	post /rs
	{
		"name": string,
		"code": string,
		"type": number,
		"parent_id": number,
		"extension": string,
	}
	例：
	{
		"name": "role23",
		"parent_id": 1111,
		"type": 2,
		"code": "xxxx33"
	}
	
	get /rs

### 1.grant permissions to role
	post /p
	{
		"role_id": number,
		"resource_id": number,
		"action": number（枚举值）
	}
	例：
	{
		"role_id": 123,
		"resource_id": 12,
		"action": 2
	}
	
	get /p?role_id={role_id}