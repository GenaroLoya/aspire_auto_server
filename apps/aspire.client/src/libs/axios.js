import axios from 'axios';
import ENV from '../../env.config';

const ax = axios.create({
  baseURL: ENV.API_URL, // Cambia esto a la URL de tu API.
  headers: {
    'Content-Type': 'application/json', // Configura el encabezado para enviar JSON.
  },
});

ax.interceptors.request.use((config) => {
  return config;
});

ax.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default ax;
