import React, { useEffect, useState } from "react";
import { getEvents } from "../api/events";
import EventCard from "../components/EventCard";

const Dashboard: React.FC = () => {
    const [events, setEvents] = useState([]);

    useEffect(() => {
        getEvents().then(setEvents);
    }, []);

    return (
        <div>
            <h2>Upcoming Events</h2>
            {events.map((event) => (
                <EventCard key={event.id} event={event} />
            ))}
        </div>
    );
};

export default Dashboard;
