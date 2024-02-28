create user api_writer WITH PASSWORD 'api_writer';
grant connect, create on database yummlog to api_writer;
grant all privileges on database yummlog to api_writer;
grant all privileges on schema yummlog to api_writer;
grant all privileges on foodposts to api_writer;
grant all privileges on postdetails to api_writer;

create user api_reader WITH PASSWORD 'api_reader';
GRANT CONNECT ON DATABASE yummlog TO api_reader;
GRANT USAGE ON SCHEMA yummlog TO api_reader;
GRANT SELECT ON foodposts TO api_reader;
GRANT SELECT ON postdetails TO api_reader;