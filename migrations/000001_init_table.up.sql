CREATE TYPE gender_type AS ENUM ('M', 'F');

CREATE TABLE hospital (
  id SERIAL PRIMARY KEY,
  code VARCHAR(20) NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT NULL
);

CREATE TABLE staff (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL,
  password VARCHAR(255) NOT NULL,
  hospital_id INT NOT NULL REFERENCES hospital(id) ON DELETE CASCADE,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT NULL,
  UNIQUE(username, hospital_id)
);

CREATE TABLE patient (
  id SERIAL PRIMARY KEY,
  patient_hn VARCHAR(20) NOT NULL,
  first_name_th VARCHAR(100),
  middle_name_th VARCHAR(100),
  last_name_th VARCHAR(100),
  first_name_en VARCHAR(100),
  middle_name_en VARCHAR(100),
  last_name_en VARCHAR(100),
  date_of_birth DATE,
  national_id CHAR(13) CHECK (char_length(national_id) = 13),
  passport_id VARCHAR(9) CHECK (char_length(passport_id) <= 9),
  phone_number VARCHAR(15) CHECK (phone_number ~ '^\+?[0-9]{9,15}$'),
  email VARCHAR(320) UNIQUE CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
  gender gender_type,
  hospital_id INT NOT NULL REFERENCES hospital(id) ON DELETE CASCADE,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT NULL,
  UNIQUE(patient_hn, hospital_id)
);

CREATE INDEX idx_patient_national_id ON patient(national_id);
CREATE INDEX idx_patient_passport_id ON patient(passport_id);
CREATE INDEX idx_patient_phone ON patient(phone_number);
CREATE INDEX idx_patient_name ON patient(first_name_th, first_name_en, last_name_th, last_name_en, middle_name_th, middle_name_en);
CREATE INDEX idx_patient_date_of_birth ON patient(date_of_birth);
CREATE INDEX idx_patient_hospital_id ON patient(hospital_id);
CREATE INDEX idx_patient_email ON patient(email);

INSERT INTO hospital (code, name)
VALUES
('HOSP-A', 'Hospital A'),
('HOSP-B', 'Hospital B');

INSERT INTO patient (
    patient_hn, first_name_th, middle_name_th, last_name_th,
    first_name_en, middle_name_en, last_name_en,
    date_of_birth, national_id, passport_id, phone_number, email,
    gender, hospital_id
)
VALUES
('HN001', 'สมชาย', NULL, 'ใจดี', 'Somchai', NULL, 'Jaidee',
 '1985-04-12', '1234567890123', NULL, '0811111111', 'somchai@example.com', 'M',
 (SELECT id FROM hospital WHERE code = 'HOSP-A')),

('HN002', 'สมหญิง', NULL, 'สวยงาม', 'Somying', NULL, 'Sungnum',
 '1990-08-20', '9876543210987', NULL, '0822222222', 'somying@example.com', 'F',
 (SELECT id FROM hospital WHERE code = 'HOSP-A')),

('HN003', 'จอน', NULL, 'โดว', 'John', NULL, 'Doe',
 '1975-01-05', NULL, 'P1234567', '0833333333', 'john.doe@example.com', 'M',
 (SELECT id FROM hospital WHERE code = 'HOSP-A')),

('HN010', 'วิชัย', NULL, 'รุ่งเรือง', 'Wichai', NULL, 'Rungrueng',
 '1982-03-22', '2223334445556', NULL, '0844444444', 'wichai@example.com', 'M',
 (SELECT id FROM hospital WHERE code = 'HOSP-B')),

('HN011', 'สุภาวดี', NULL, 'อารีย์', 'Supawadee', NULL, 'Aree',
 '1995-11-15', '3334445556667', NULL, '0855555555', 'supawadee@example.com', 'F',
 (SELECT id FROM hospital WHERE code = 'HOSP-B')),

('HN012', 'อลิส', NULL, 'สมิท', 'Alice', NULL, 'Smith',
 '1988-07-30', NULL, 'P7654321', '0866666666', 'alice.smith@example.com', 'F',
 (SELECT id FROM hospital WHERE code = 'HOSP-B'));

