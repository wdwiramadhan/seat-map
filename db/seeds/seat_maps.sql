INSERT INTO
    seat_maps (id, aircraft)
VALUES
    (
        '01979faf-267b-7663-9ac3-b2ca127cd683',
        'Boeing 737'
    );

INSERT INTO
    cabins (
        id,
        seat_map_id,
        deck,
        seat_columns,
        first_row,
        last_row
    )
VALUES
    (
        '01979faf-267b-7663-9ac3-b2ca127cd684',
        '01979faf-267b-7663-9ac3-b2ca127cd683',
        'Main',
        ARRAY['LEFT_SIDE', 'A', 'B', 'C', 'AISLE', 'D','E', 'F', 'RIGHT_SIDE'],
        1,
        2
    );

INSERT INTO
    seat_rows (id, cabin_id, row_number)
VALUES
    (
        '01979faf-267b-7663-9ac3-b2ca127cd685',
        '01979faf-267b-7663-9ac3-b2ca127cd684',
        0
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd686',
        '01979faf-267b-7663-9ac3-b2ca127cd684',
        1
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd687',
        '01979faf-267b-7663-9ac3-b2ca127cd684',
        2
    );

INSERT INTO
    seats (
        id,
        seat_row_id,
        storefront_slot_code,
        available,
        code
    )
VALUES
    (
        '01979faf-267b-7663-9ac3-b2ca127cd688',
        '01979faf-267b-7663-9ac3-b2ca127cd685',
        'BLANK',
        false,
        ''
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd689',
        '01979faf-267b-7663-9ac3-b2ca127cd685',
        'BULKHEAD',
        false,
        ''
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd690',
        '01979faf-267b-7663-9ac3-b2ca127cd685',
        'BULKHEAD',
        false,
        ''
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd691',
        '01979faf-267b-7663-9ac3-b2ca127cd685',
        'BULKHEAD',
        false,
        ''
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd692',
        '01979faf-267b-7663-9ac3-b2ca127cd685',
        'BLANK',
        false,
        ''
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd693',
        '01979faf-267b-7663-9ac3-b2ca127cd685',
        'BULKHEAD',
        false,
        ''
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd694',
        '01979faf-267b-7663-9ac3-b2ca127cd685',
        'BULKHEAD',
        false,
        ''
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd695',
        '01979faf-267b-7663-9ac3-b2ca127cd685',
        'BULKHEAD',
        false,
        ''
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd697',
        '01979faf-267b-7663-9ac3-b2ca127cd685',
        'BLANK',
        false,
        ''
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd698',
        '01979faf-267b-7663-9ac3-b2ca127cd686',
        'BLANK',
        false,
        ''
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd699',
        '01979faf-267b-7663-9ac3-b2ca127cd686',
        'SEAT',
        true,
        '1A'
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd700',
        '01979faf-267b-7663-9ac3-b2ca127cd686',
        'SEAT',
        true,
        '2A'
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd701',
        '01979faf-267b-7663-9ac3-b2ca127cd686',
        'SEAT',
        true,
        '3A'
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd702',
        '01979faf-267b-7663-9ac3-b2ca127cd686',
        'BLANK',
        false,
        ''
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd703',
        '01979faf-267b-7663-9ac3-b2ca127cd686',
        'SEAT',
        true,
        '4A'
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd704',
        '01979faf-267b-7663-9ac3-b2ca127cd686',
        'SEAT',
        true,
        '5A'
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd705',
        '01979faf-267b-7663-9ac3-b2ca127cd686',
        'SEAT',
        true,
        '6A'
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd706',
        '01979faf-267b-7663-9ac3-b2ca127cd686',
        'BLANK',
        false,
        ''
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd707',
        '01979faf-267b-7663-9ac3-b2ca127cd687',
        'BLANK',
        false,
        ''
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd708',
        '01979faf-267b-7663-9ac3-b2ca127cd687',
        'SEAT',
        true,
        '1B'
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd709',
        '01979faf-267b-7663-9ac3-b2ca127cd687',
        'SEAT',
        true,
        '2B'
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd710',
        '01979faf-267b-7663-9ac3-b2ca127cd687',
        'SEAT',
        true,
        '3B'
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd711',
        '01979faf-267b-7663-9ac3-b2ca127cd687',
        'BLANK',
        false,
        ''
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd712',
        '01979faf-267b-7663-9ac3-b2ca127cd687',
        'SEAT',
        true,
        '4B'
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd713',
        '01979faf-267b-7663-9ac3-b2ca127cd687',
        'SEAT',
        true,
        '5B'
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd714',
        '01979faf-267b-7663-9ac3-b2ca127cd687',
        'SEAT',
        true,
        '6B'
    ),
    (
        '01979faf-267b-7663-9ac3-b2ca127cd715',
        '01979faf-267b-7663-9ac3-b2ca127cd687',
        'BLANK',
        false,
        ''
    );
