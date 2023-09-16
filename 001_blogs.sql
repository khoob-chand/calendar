-- Write your migrate up statements here
-- drop table blogs;
CREATE TABLE blogs(
   id SERIAL PRIMARY KEY,
   title text,
   summary text,
   content text ,   
   create_at timestamp default current_timestamp,
   published_at text);





