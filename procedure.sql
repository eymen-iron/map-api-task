DELIMITER $$

DROP PROCEDURE IF EXISTS `GetNearestLocationsWithPagination`$$
CREATE DEFINER=`root`@`localhost` PROCEDURE `GetNearestLocationsWithPagination` (
    IN `limit_count` INT,
    IN `offset_count` INT,
    IN `input_lat` FLOAT,
    IN `input_long` FLOAT
)   
/// Haversine formula
BEGIN    
    SELECT id, name, latitude, longitude, marker,
           6371 * 2 * ASIN(SQRT(
               POWER(SIN((RADIANS(input_lat) - RADIANS(latitude)) / 2), 2) +
               COS(RADIANS(input_lat)) * COS(RADIANS(latitude)) *
               POWER(SIN((RADIANS(input_long) - RADIANS(longitude)) / 2), 2)
           )) AS distance
    FROM locations
    ORDER BY distance
    LIMIT limit_count OFFSET offset_count;
END$$

DELIMITER ;
