/**
 * @typedef {Object} Response
 * @template T
 * @property {number} code
 * @property {string} message
 * @property {T} data
 */

/**
 * @typedef {Object} ListData
 * @template T
 * @property {number} total
 * @property {T[]} data
 */

/**
 * @template T
 * @typedef {Response<ListData<T>>} ListResponse
 */

export class MediaClient {
  /**
   *
   * @param {import("axios").AxiosInstance} axios
   */
  constructor(axios) {
    this.axios = axios;
  }

  /**
   * @typedef {Object} Media
   * @property {number} id
   * @property {string} created_at
   * @property {string} updated_at
   * @property {string} path
   * @property {string} name
   * @property {number} media_type
   * @property {number} information_id
   * @property {number} cover_id
   */

  /**
   * 列出媒体
   * @param {string|null} q
   * @param {number} page
   * @param {number} page_size
   * @returns {Promise<ListResponse<Media>>}
   */
  async list(q, page, page_size) {
    const payload = { q, page, page_size };
    return (await this.axios.post("/api/v1/media/list", payload)).data;
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

export class PathClient {
  /**
   * @param {import("axios").AxiosInstance} axios
   */
  constructor(axios) {
    this.axios = axios;
  }

  /**
   * @typedef {Object} PathEntry
   * @property {string} name
   * @property {boolean} is_dir
   * @property {string} path
   * @property {string} updated_at
   */

  /**
   *
   * @param {number} page
   * @param {number} page_size
   * @returns {Promise<ListResponse<PathEntry>>}
   */
  async list(page, page_size) {
    const payload = { page, page_size };
    return (await this.axios.post("/api/v1/path/list", payload)).data;
  }
}

export class AnimeClient {
  /**
   * @param {import("axios").AxiosInstance} axios
   */
  constructor(axios) {
    this.axios = axios;
  }

  /**
   * @typedef {Object} AnimeDTO
   * @property {uint} id
   * @property {string} created_at
   * @property {string} updated_at
   * @property {string} title
   * @property {string[]} synonyms
   * @property {number} total_episodes
   * @property {number} release_year
   * @property {string} season
   * @property {string} status
   * @property {string[]} tags
   */

  /**
   * 列出动漫
   * @param {number} page
   * @param {number} page_size
   * @returns {Promise<ListResponse<AnimeDTO>>}
   */
  async list(page, page_size) {
    const payload = { page, page_size };
    return (await this.axios.post("/api/v1/anime/list", payload)).data;
  }

  /**
   * @typedef {Object} AnimeOfflineDatabaseDTO
   * @property {number} id
   * @property {string} title
   * @property {string[]} sources
   * @property {string} type
   * @property {number} episodes
   * @property {string} status
   * @property {number} year
   * @property {string} season
   * @property {string} picture
   * @property {string} thumbnail
   * @property {number} duration
   * @property {string[]} synonyms
   * @property {string[]} tags
   */

  /**
   * 搜索动漫
   * @param {String} term
   * @param {number} page
   * @param {number} page_size
   * @returns {Promise<ListResponse<AnimeOfflineDatabaseDTO>>}
   */
  async search(term, page, page_size) {
    const payload = { term, page, page_size };
    return (await this.axios.post("/api/v1/anime/search", payload)).data;
  }

  /**
   * 获取动画信息
   * @param {number} id
   * @returns {Promise<Response<AnimeOfflineDatabaseDTO>>}
   */
  async info(id) {
    return (await this.axios.get(`/api/v1/anime/info/${id}`)).data;
  }
}
