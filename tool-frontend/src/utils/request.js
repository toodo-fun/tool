import axios from "axios";

const service = axios.create({
    baseURL: "http://127.0.0.1:30612/tool/server",
    timeout: 1000 * 60 * 60
})

service.interceptors.request.use(
    config => {
        config.headers["X-Token"] = "token"
        return config
    },
    error => {
        console.log(error);
        return Promise.reject(error)
    },
)

service.interceptors.response.use(
    response => {
        const res = response.data
        if (!res.code) {
            return res
        }
        if (res.code !== "00000") {
            console.log("请求出错", res.message);
            return Promise.reject(new Error(res.message || '请求出错'))
        } else {
            return res.result
        }
    },
    error => {
        console.log("请求出错", error);
        return Promise.reject(error)
    },
)

export default service
