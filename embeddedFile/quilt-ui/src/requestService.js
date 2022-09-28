import axios from 'axios';

const get = async (url, config) => axios.get(url, config);
const post = async (url, data, config) => axios.post(url, data, config);
const put = async (url, data, config) => axios.put(url, data, config);

const RequestService = {
  get,
  post,
  put,
};

export default RequestService;
