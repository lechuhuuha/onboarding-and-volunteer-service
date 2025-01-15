-- +goose Up
-- +goose NO TRANSACTION
-- Insert roles
INSERT INTO
    roles (id, name)
VALUES
    (1, 'Guest'),
    (2, 'Applicant'),
    (3, 'Volunteer'),
    (4, 'Admin');

-- Insert departments
INSERT INTO
    departments (id, name, address, status)
VALUES
    (1, 'Health Services', '123 Main St, City A', 1),
    (2, 'Education Support', '456 Elm St, City B', 1),
    (3, 'Disaster Relief', '789 Oak St, City C', 1);

-- Insert countries
INSERT INTO
    countries (id, name, status)
VALUES
    (1, 'Country A', 1),
    (2, 'Country B', 1),
    (3, 'Country C', 1);

-- Insert a sample guest user
INSERT INTO
    users (
        id,
        role_id,
        email,
        password,
        name,
        surname,
        gender,
        dob,
        mobile,
        country_id,
        resident_country_id,
        verification_status,
        status
    )
VALUES
    (
        1,
        1,
        'guest@example.com',
        'hashed_password',
        'Guest',
        'User',
        'Other',
        '2000-01-01',
        '1234567890',
        1,
        1,
        0,
        1
    );

-- Insert an applicant user
INSERT INTO
    users (
        id,
        role_id,
        email,
        password,
        name,
        surname,
        gender,
        dob,
        mobile,
        country_id,
        resident_country_id,
        verification_status,
        status
    )
VALUES
    (
        2,
        2,
        'applicant@example.com',
        'hashed_password',
        'Applicant',
        'User',
        'Female',
        '1995-06-15',
        '0987654321',
        2,
        2,
        0,
        1
    );

-- Insert a volunteer user
INSERT INTO
    users (
        id,
        role_id,
        email,
        password,
        name,
        surname,
        gender,
        dob,
        mobile,
        country_id,
        resident_country_id,
        verification_status,
        status
    )
VALUES
    (
        3,
        3,
        'volunteer@example.com',
        'hashed_password',
        'Volunteer',
        'User',
        'Male',
        '1990-05-20',
        '1231231234',
        3,
        3,
        1,
        1
    );

-- Insert an admin user
INSERT INTO
    users (
        id,
        role_id,
        email,
        password,
        name,
        surname,
        gender,
        dob,
        mobile,
        country_id,
        resident_country_id,
        verification_status,
        status
    )
VALUES
    (
        4,
        4,
        'admin@example.com',
        'hashed_password',
        'Admin',
        'User',
        'Other',
        '1985-11-11',
        '3213213214',
        1,
        1,
        1,
        1
    );

-- Insert volunteer details
INSERT INTO
    volunteer_details (id, user_id, department_id, status)
VALUES
    (1, 3, 1, 1);

-- Insert requests (example: Applicant registration request)
INSERT INTO
    requests (
        id,
        user_id,
        type,
        status,
        reject_notes,
        verifier_id
    )
VALUES
    (1, 2, 'Registration', 0, NULL, NULL);

-- Insert user identities
INSERT INTO
    user_identities (
        id,
        user_id,
        number,
        type,
        status,
        expiry_date,
        place_issued
    )
VALUES
    (
        1,
        3,
        'ID123456',
        'Passport',
        1,
        '2030-12-31',
        'City A'
    );