import React from "react";

const EventCard: React.FC<{ event: any }> = ({ event }) => {
    return (
        <div>
            <h3>{event.name}</h3>
            <p>{event.date}</p>
        </div>
    );
};

export default EventCard;
