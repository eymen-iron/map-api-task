DELIMITER //

CREATE PROCEDURE GetNearestLocationsWithPagination(
    IN limit_count INT,
    IN offset_count INT,
    IN input_lat FLOAT,
    IN input_long FLOAT
)
BEGIN    
    SELECT id, name, latitude, longitude, marker,
           (6371 * acos(cos(radians(input_lat)) * cos(radians(latitude)) * cos(radians(longitude) - radians(input_long)) + sin(radians(input_lat)) * sin(radians(latitude)))) AS distance
    FROM locations
    ORDER BY distance
    LIMIT limit_count OFFSET offset_count;
END //

DELIMITER ;
