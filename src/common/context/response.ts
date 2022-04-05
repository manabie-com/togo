export interface ResponseInterface<D = any> {
  status?: 'success' | 'error' | 'pending';
  message?: string;
  data?: D;
}
