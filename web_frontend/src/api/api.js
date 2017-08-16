import axios from 'axios';

var instance = axios.create({
  baseURL: '/api',
  timeout: 2500,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  },

});

instance.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    var user = sessionStorage.getItem('user');
    if(user) {
      user = JSON.parse(user);
      config.headers.token = user.token;
    }

    console.log(config);
    return config;
  }, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  });

export const requestLogin = params => { return instance.post(`/login`, params).then(res => res.data); };

export const register = params => { return instance.post(`/register`, params).then(res => res.data); };
export const validateEmail = params => { return instance.post(`/validate`, params).then(res => res.data); };
export const sendMail = params => { return instance.post(`/mail`, params).then(res => res.data); };


export const getUserListPage = params => { return instance.get(`/user/`, {params: params}).then(res => res.data); };

export const getUser = params => {return instance.get(`/user/${params.id}`).then(res => res.data);}
export const editUser = params => {return instance.put(`/user/${params.id}`, params).then(res => res.data);}

export const getArticleListPage = params => { return instance.get(`/article/`, {params: params}).then(res => res.data); };

export const addArticle = params => { return instance.post(`/article/`, params).then(res => res.data); };
export const getArticle = params => {return instance.get(`/article/${params.id}`).then(res => res.data);}

export const getCommentListPage = params => { return instance.get(`/comment/`, {params: params}).then(res => res.data); };
export const addComment = params => { return instance.post(`/comment/`, params).then(res => res.data); };
