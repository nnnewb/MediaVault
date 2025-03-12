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
     * @param {import('axios').AxiosInstance} axios
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
     * 列出媒体
     * @param {number} page
     * @param {number} page_size
     * @returns {Promise<ListResponse<PathEntry>>}
     */
    async list(page, page_size) {
        const payload = {page, page_size};
        return (await this.axios.post("/api/v1/media/list", payload)).data;
    }

    async add(paths) {
        const payload = {paths};
        return await this.axios.post("/api/v1/media/add", payload);
    }

    async scan(paths) {
        const payload = {paths};
        return await this.axios.post("/api/v1/media/scan", payload);
    }
}

export class TaskClient {
    constructor(axios) {
        this.axios = axios;
    }

    async list(page, page_size) {
        const payload = {page, page_size};
        return await this.axios.post("/api/v1/task/list", payload);
    }
}

export class AnimeClient {
    /**
     * @param {import('axios').AxiosInstance} axios
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
        const payload = {page, page_size};
        return (await this.axios.post("/api/v1/anime/list", payload)).data;
    }
}
