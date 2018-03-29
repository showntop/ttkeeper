# ttkeeper application interface document

- devp host: localhost:8089
- test host: localhost:8089
- prod host: localhost:8089

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