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
	get /u