import axios from "axios";
import { getToken } from "../utils/auth";

const API_URL = "http://localhost:8080/events";

export const getEvents = async () => {
    const response = await axios.get(API_URL);
    return response.data;
};

export const createEvent = async (eventData: any) => {
    const response = await axios.post(API_URL, eventData, {
        headers: { Authorization: `Bearer ${getToken()}` }
    });
    return response.data;
};
