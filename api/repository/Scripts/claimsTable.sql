CREATE TABLE claims (
	clientid  		varchar(305),
	clientpassword  varchar(305),
	scope			varchar(50),
	userid			int UNIQUE,
    PRIMARY KEY(userid)
);