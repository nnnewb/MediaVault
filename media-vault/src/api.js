export class MediaClient {
  constructor(axios) {
    this.axios = axios;
  }

  async list(page, page_size) {
    const payload = { page, page_size };
    return await this.axios.post("/api/v1/media/list", payload);
  }

  async add(paths) {
    const payload = { paths };
    return await this.axios.post("/api/v1/media/add", payload);
  }

  async scan(paths) {
    const payload = { paths };
    return await this.axios.post("/api/v1/media/scan", payload);
  }
}

export class TaskClient {
  constructor(axios) {
    this.axios = axios;
  }

  async list(page, page_size) {
    const payload = { page, page_size };
    return await this.axios.post("/api/v1/task/list", payload);
  }
}
