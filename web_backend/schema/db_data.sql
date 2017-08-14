insert into users(username,email,password,type,created_at)
  values('admin','774601526@qq.com','eed8c91348684d6087fd4f6f94eb0418',3,1501383651);

insert into users_info(user_id,created_at) values(1,1501383651);

  insert into configs(name,value,description,created_at)
    values("mail_user","xxx@xx.com","邮箱账号",1501383651),
    ("mail_pass","pass","邮箱密码",1501383651),
    ("mail_host","host:port","邮箱服务器地址",1501383651);
