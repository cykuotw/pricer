/**
 * Configuration file for API endpoints.
 * This file constructs the base API URL using environment variables.
 *
 * Features:
 * - Reads the server URL and API path from environment variables.
 * - Exports the constructed API URL for use throughout the application.
 *
 * Dependencies:
 * - Environment variables: `VITE_SERVER_URL`, `VITE_API_PATH`
 *
 * @module config
 */

const SERVER_URL = import.meta.env.VITE_SERVER_URL;
const API_PATH = import.meta.env.VITE_API_PATH;

export const API_URL = `http://${SERVER_URL}${API_PATH}`;
