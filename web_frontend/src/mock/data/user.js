import Mock from 'mockjs';
const LoginUsers = [
  {
    id: 1,
    username: 'user',
    password: '123456',
    avatar: 'https://raw.githubusercontent.com/taylorchen709/markdown-images/master/vueadmin/user.png',
    name: '张某某'
  }
];

const Users = [];

for (let i = 0; i < 10; i++) {
  var bgs = [
    'https://cdn.shopify.com/s/files/1/0691/5403/t/139/assets/insta-2.jpg?12817922976150776315',
    'https://cdn.shopify.com/s/files/1/0691/5403/t/139/assets/insta-1.jpg?12817922976150776315',
    'https://cdn.shopify.com/s/files/1/0691/5403/t/139/assets/insta-3.jpg?12817922976150776315'
  ];
  Users.push(Mock.mock({
    id: Mock.Random.guid(),
    name: Mock.Random.cname(),
    bg:bgs[Mock.Random.integer(0,2)],
    avatar:Mock.Random.dataImage('100x100','head'),
    about: Mock.Random.ctitle(3,10),
    descripttion: Mock.Random.cword(5,30),

  }));
}

export { LoginUsers, Users };
