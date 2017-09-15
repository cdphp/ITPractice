import axios from 'axios';

var instance = axios.create({
  baseURL: '/api',
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  },

});

instance.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    var user = localStorage.getItem('user');
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
export const upload = params => { return instance.post(`/upload`, params).then(res => res.data); };
export const forget = params => { return instance.post(`/forget`, params).then(res => res.data); };

export const oauthGithub = params => {return instance.get(`/oauth`,{params: params}).then(res => res.data);}
export const getGithubUser = params => {return axios.get(`https://api.github.com/user`,{params: params}).then(res => res.data);}


export const getUserListPage = params => { return instance.get(`/user/`, {params: params}).then(res => res.data); };

export const getUser = params => {return instance.get(`/user/${params.id}`).then(res => res.data);}
export const editUser = params => {return instance.put(`/user/${params.id}`, params).then(res => res.data);}
export const resetPass = params => {return instance.post(`/user/resetPass`, params).then(res => res.data);}

export const getArticleListPage = params => { return instance.get(`/article/`, {params: params}).then(res => res.data); };

export const addArticle = params => { return instance.post(`/article/`, params).then(res => res.data); };
export const getArticle = params => {return instance.get(`/article/${params.id}`).then(res => res.data);}
export const delArticle = params => { return instance.delete(`/article/${params.id}`).then(res => res.data); };
export const editArticle = params => {return instance.put(`/article/${params.id}`, params).then(res => res.data);}

export const getCommentListPage = params => { return instance.get(`/comment/`, {params: params}).then(res => res.data); };
export const addComment = params => { return instance.post(`/comment/`, params).then(res => res.data); };

export const addRelation = params => { return instance.post(`/relation/`, params).then(res => res.data); };
export const getRelationListPage = params => { return instance.get(`/relation/`, {params: params}).then(res => res.data); };

export const getCompanyListPage = params => { return instance.get(`/company/`, {params: params}).then(res => res.data); };
export const addCompany = params => { return instance.post(`/company/`, params).then(res => res.data); };

export const getQuestionListPage = params => { return instance.get(`/question/`, {params: params}).then(res => res.data); };
export const addQuestion = params => { return instance.post(`/question/`, params).then(res => res.data); };
export const getQuestion = params => {return instance.get(`/question/${params.id}`).then(res => res.data);}
export const delQuestion = params => { return instance.delete(`/question/${params.id}`).then(res => res.data); };
export const editQuestion = params => {return instance.put(`/question/${params.id}`, params).then(res => res.data);}

export const getAnswerListPage = params => { return instance.get(`/answer/`, {params: params}).then(res => res.data); };
export const addAnswer = params => { return instance.post(`/answer/`, params).then(res => res.data); };
export const evaluteAnswer = params => { return instance.post(`/answer/evalute`, params).then(res => res.data); };

export const getMessageListPage = params => { return instance.get(`/message/`, {params: params}).then(res => res.data); };
export const addMessage = params => { return instance.post(`/message/`, params).then(res => res.data); };
