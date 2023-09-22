import axios from 'axios';

class CoursService {
    getAll(): Promise<any> {
      return axios.get("http://localhost:4000/cours");
    }
  
    get(id: any): Promise<any> {
      return axios.get(`http://localhost:4000/cours/${id}`);
    }
  
    // create(data: any): Promise<any> {
    //   return axios.post("http://localhost:4000/cours", data);
    // }
  
    // update(id: any, data: any): Promise<any> {
    //   return axios.put(`http://localhost:4000/cours/${id}`, data);
    // }
  
    // delete(id: any): Promise<any> {
    //   return axios.delete(`http://localhost:4000/cours/${id}`);
    // }
  }
  
  export default new CoursService();