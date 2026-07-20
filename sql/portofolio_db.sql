--
-- PostgreSQL database dump
--

-- Dumped from database version 14.7
-- Dumped by pg_dump version 14.7

-- Started on 2026-05-30 00:53:51

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET search_path = public;
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 209 (class 1259 OID 21360)
-- Name: admins; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.admins (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(255) NOT NULL,
    username character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- TOC entry 212 (class 1259 OID 21394)
-- Name: certificates; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.certificates (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    title character varying(255) NOT NULL,
    issuer character varying(255) NOT NULL,
    issue_date character varying(50) NOT NULL,
    credential_url character varying(255),
    thumbnail character varying(255),
    "order" bigint DEFAULT 0,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- TOC entry 211 (class 1259 OID 21383)
-- Name: experiences; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.experiences (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company character varying(255) NOT NULL,
    role character varying(255) NOT NULL,
    location character varying(255),
    start_date character varying(50) NOT NULL,
    end_date character varying(50),
    current_job boolean DEFAULT false,
    description text NOT NULL,
    "order" bigint DEFAULT 0,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- TOC entry 214 (class 1259 OID 21414)
-- Name: profiles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.profiles (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(100) NOT NULL,
    title character varying(255) NOT NULL,
    subtitle character varying(255) NOT NULL,
    intro text,
    bio text,
    avatar character varying(255),
    resume_url character varying(255),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    github character varying(255),
    instagram character varying(255),
    linkedin character varying(255),
    email character varying(255)
);


--
-- TOC entry 210 (class 1259 OID 21373)
-- Name: projects; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.projects (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    title character varying(255) NOT NULL,
    description text NOT NULL,
    thumbnail character varying(255),
    github_url character varying(255),
    demo_url character varying(255),
    tech_stack character varying(255),
    "order" bigint DEFAULT 0,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- TOC entry 213 (class 1259 OID 21404)
-- Name: skills; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.skills (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(255) NOT NULL,
    category character varying(255) NOT NULL,
    "order" bigint DEFAULT 0,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- TOC entry 3355 (class 0 OID 21360)
-- Dependencies: 209
-- Data for Name: admins; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.admins (id, name, username, email, password, created_at, updated_at, deleted_at) FROM stdin;
de22f31d-00e5-4702-91dc-afe720f84106	Portfolio Admin	admin	admin@example.com	$2a$10$eqFZCWgGhb46MGlsxueUhu4ioMv.GrstE2xAwvTXF0dv3YOTjVcOq	2026-05-28 23:54:33.608824+07	2026-05-28 23:54:33.608824+07	\N
\.


--
-- TOC entry 3358 (class 0 OID 21394)
-- Dependencies: 212
-- Data for Name: certificates; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.certificates (id, title, issuer, issue_date, credential_url, thumbnail, "order", created_at, updated_at, deleted_at) FROM stdin;
28422028-6f57-421c-8bf9-c3d1a038d80d	Google Data Analytics Professional Certificate	Coursera / Google	May 2025	https://coursera.org/verify/example	/uploads/certificates/google-data.png	1	2026-05-29 00:06:43.53397+07	2026-05-29 00:06:43.53397+07	2026-05-30 00:18:23.005473+07
cc6d4a7d-d66b-457f-a4f3-bc1f9a65125f	Code Generation and Optimization Using IBM Granite	IBM	September 2025			1	2026-05-30 00:19:21.775132+07	2026-05-30 00:19:21.775132+07	\N
3b7a2d5d-80fc-4ec8-8e5f-3bce767ee3aa	Junior Mobile Programmer (Vocational School Graduate Academy)	Digital Talent Scholarship	August 2025			2	2026-05-30 00:19:46.107979+07	2026-05-30 00:19:46.107979+07	\N
819cf83f-edc5-4f1c-a418-93b374b03697	Intro to Software Engineering	RevoU	August 2025			3	2026-05-30 00:20:01.462142+07	2026-05-30 00:20:01.462142+07	\N
724747d8-3c9d-47f2-b9f6-4a1d7c055bf4	Belajar Dasar Visualisasi Data	Dicoding Indonesia	September 2025			4	2026-05-30 00:20:22.255988+07	2026-05-30 00:20:22.255988+07	\N
023d06d9-d324-4386-9d89-94052869bcab	Mindset Digital (Micro Skill)	Digital Talent Scholarship	August 2025			5	2026-05-30 00:20:37.060671+07	2026-05-30 00:20:37.060671+07	\N
a8237bb0-f9e3-4aa3-93d1-d694f0ee7830	Intro to Data Analytics	RevoU	March 2024			6	2026-05-30 00:20:50.562805+07	2026-05-30 00:20:50.562805+07	\N
\.


--
-- TOC entry 3357 (class 0 OID 21383)
-- Dependencies: 211
-- Data for Name: experiences; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.experiences (id, company, role, location, start_date, end_date, current_job, description, "order", created_at, updated_at, deleted_at) FROM stdin;
e5ad3d8e-4f3e-4330-b999-8f91d0a40797	AEON Mall	Data Analyst & Backend Intern	Jakarta, ID	July 2024	December 2024	f	- Developed automated supply chain forecasting dashboards...\n- Maintained database consistency.	1	2026-05-29 00:06:27.189498+07	2026-05-29 00:06:27.189498+07	2026-05-30 00:12:11.126532+07
17232912-90a6-4cb5-8f50-53b652561439	Celerates	Data Science	Jakarta Selatan	September 2024	December 2024	f	- MSIB Batch 7	2	2026-05-29 02:14:00.898148+07	2026-05-29 02:14:00.898148+07	2026-05-30 00:12:15.213116+07
43173b56-7d24-4417-b67e-0fce02624e69	Madani Kreatif (CV. Madani Berkah Abadi)	Freelance Fullstack Developer	Remote	January 2026	May 2026	f	- Developed and maintained fullstack web applications using React, Golang Fiber, PostgreSQL, and REST API architecture.\n- Designed and implemented secure authentication and role-based access control (RBAC) systems.\n- Built backend services for content management, order processing, and file upload functionality.\n- Developed responsive and user-friendly interfaces aligned with business requirements.\n- Integrated frontend and backend services to ensure efficient data flow and system reliability.	1	2026-05-30 00:14:28.462596+07	2026-05-30 00:14:28.462596+07	\N
27203161-4ebf-4314-8c70-576adb03cf2b	Kementerian Perumahan dan Kawasan Permukiman (PKP)	Backend Developer Intern	Yogyakarta, Indonesia	October 2025	December 2025	f	- Developed RESTful API services using Node.js, Express.js, and MySQL.\n- Designed database schemas and optimized queries for application performance.\n- Implemented CRUD operations and backend business logic for government service applications.\n- Collaborated with frontend developers to integrate APIs and ensure seamless functionality.\n- Participated in testing, debugging, and deployment activities during development cycles.	2	2026-05-30 00:15:43.624246+07	2026-05-30 00:15:43.624246+07	\N
bb69f86b-57e0-40c5-97a5-daf890eb3744	Celerates (PT. Mitra Talenta Grup)	Data Science Mentee	Remote	September 2024	December 2024	f	- Developed a deep learning-based sentiment analysis model using LSTM architecture for social media text classification.\n- Performed NLP preprocessing including text cleaning, tokenization, stopword removal, TF-IDF vectorization, and sequence padding.\n- Integrated Gemini API for automated topic extraction and text analysis.\n- Built an interactive Streamlit application for real-time sentiment prediction and topic modeling.\n- Evaluated model performance using standard machine learning and NLP metrics.	3	2026-05-30 00:16:46.218235+07	2026-05-30 00:16:46.218235+07	\N
\.


--
-- TOC entry 3360 (class 0 OID 21414)
-- Dependencies: 214
-- Data for Name: profiles; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.profiles (id, name, title, subtitle, intro, bio, avatar, resume_url, created_at, updated_at, deleted_at, github, instagram, linkedin, email) FROM stdin;
025605e7-32cb-4820-bb02-598643904637	Shah Delphi Muhammad	Shah Delphi.	I build things for the web.	Final Year Informatics Student | AI Enthusiast | Data Analytics & Supply Chain Analytics | Software Developer	Final year Informatics student with strong interest in Artificial Intelligence, Data Analytics, Supply Chain Analytics, and Software Development. Experienced in developing AI-based and data-driven solutions through academic and independent projects involving computer vision, sentiment analysis, forecasting systems, backend development, and mobile applications.\nExperienced in building machine learning models, forecasting systems, backend services, and fullstack applications for academic and real-world use cases. Worked on projects related to YOLOv8 segmentation, NLP-based sentiment analysis, and supply chain forecasting using SARIMAX and LSTM models.\nInterested in AI engineering, data analytics, forecasting, and technology-driven business solutions. Passionate about continuous learning and building impactful digital products.	/uploads/profile/2ce4e790-daec-4694-a03b-fb647a841c77-1779995300.jpg	/uploads/profile/e29434b8-b5f7-4872-aec2-ab42cc19e135-1779995308.pdf	2026-05-29 02:04:09.307901+07	2026-05-29 23:09:39.137613+07	\N	https://github.com/ShahDelphi	https://www.instagram.com/shah.delphi/	https://www.linkedin.com/in/shah-delphi-muhammad/	shahdelphimuhammad@gmail.com
\.


--
-- TOC entry 3356 (class 0 OID 21373)
-- Dependencies: 210
-- Data for Name: projects; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.projects (id, title, description, thumbnail, github_url, demo_url, tech_stack, "order", created_at, updated_at, deleted_at) FROM stdin;
42e436f4-f993-4c7d-b22d-715bd556dd77	Backend & API Projects Collection	A compilation of clean-architecture microservices and RESTful API endpoints built using Golang (Fiber/Gin) and Node.js. Features rate limiting, structured logging, JWT authentication, and comprehensive unit testing.	/uploads/projects/sample-backend.webp	https://github.com/username/backend-api-collection		Golang,Fiber,Redis,PostgreSQL,JWT,Docker	5	2026-05-28 23:54:33.614934+07	2026-05-28 23:54:33.614934+07	2026-05-29 00:06:05.149751+07
e213a389-ebd8-4d35-b6e4-b5662819d952	Portofolio	Membuat Web Portofolio	/uploads/projects/48dd8db2-0943-4eb0-88dd-8c9d24c25e0c-1779992788.jpg	https://github.com/ShahDelphi/React-Portfolio-Template-v1.git		React	0	2026-05-29 01:26:32.438147+07	2026-05-29 01:26:32.438147+07	2026-05-29 01:51:38.927245+07
68412af3-a1b6-4f57-b4ce-ea2497bb0a48	Oil Market Forecasting for Supply Chain Decision Making	Built a predictive analytics dashboard utilizing historical prices, geopolitical events, and freight capacity data to model global crude oil pricing. Provided actionable supply chain recommendations to minimize procurement costs.	/uploads/projects/sample-supplychain.webp	https://github.com/username/oil-supply-chain		Python,Pandas,XGBoost,FastAPI,React,Docker	3	2026-05-28 23:54:33.608824+07	2026-05-28 23:54:33.608824+07	2026-05-29 23:59:40.675739+07
9bdc37a4-3316-4fd3-9f62-65b2308cd591	Fullstack Printing Web App	A modern web portal for handling online printing requests. Users can upload files, choose paper options, view instant quotes, and pay online. Admins can manage incoming jobs, update printing status, and handle billing.	/uploads/projects/sample-printing.webp	https://github.com/username/printing-web-app	https://printing-demo.example.com	React,Tailwind CSS,Node.js,Express,PostgreSQL	4	2026-05-28 23:54:33.614121+07	2026-05-28 23:54:33.614121+07	2026-05-29 23:59:42.745882+07
d64801b8-0897-4487-b5f0-628e24652ac6	AEON Mall Forecasting & Ordering System	Developed an inventory forecasting engine integrated with a smart ordering workflow for AEON Mall fresh food sections. Utilized seasonal ARIMA models to reduce food waste by 22% while maintaining optimal stock availability.	/uploads/projects/755b8a00-c590-4e6e-ab48-ef73922b15c1-1779994361.png	https://github.com/username/aeon-forecasting		Golang,React,PostgreSQL,Python,ARIMA	2	2026-05-28 23:54:33.608824+07	2026-05-29 01:52:42.64092+07	2026-05-29 23:59:44.564579+07
00c81f7e-22e8-48cd-a653-84df16eed09c	Updated YOLOv8 Palm Oil Project	Updated description text...	/uploads/projects/8d650f5a-edc3-438a-821d-b3054d0f2ecf-1779994173.png	https://github.com/example/yolo	https://demo.example.com	Python,YOLOv8,PyTorch,OpenCV,Streamlit	1	2026-05-28 23:54:33.608824+07	2026-05-29 01:49:34.358112+07	2026-05-29 23:59:46.566893+07
168d8c5c-3c53-48b0-92d0-9f04115c5036	Simulation-Based Analysis for Sales Forecasting, Ordering Policy, and Inventory Optimization of Perishable Products in Retail Industry	Developed a simulation-based analytics project for sales forecasting, inventory control, and ordering policy optimization of perishable retail products.\n\nImplemented SARIMAX, LSTM, and hybrid SARIMAX-LSTM models to analyze seasonal demand patterns and promotional spikes while improving stock management efficiency and minimizing product waste.	/uploads/projects/eeb9d2b1-943f-44fb-85e0-9273d3cb2cdd-1780074171.png			Supply Chain Analytics, Time Series Forecasting, Long Short-term Memory (LSTM),SARIMAX, Big Data Analytics, Python (Bahasa Pemrograman)	1	2026-05-30 00:02:52.949473+07	2026-05-30 00:02:52.949473+07	\N
eaf178bc-3c14-4d33-af9b-02c597c0be3f	YOLOv8 Segmentation for Palm Oil Leaf Disease Detection	Developed a computer vision model for palm oil leaf disease severity detection using YOLOv8 segmentation.\n\nPerformed dataset annotation, model training, validation, and evaluation using segmentation metrics to improve disease visualization and detection accuracy.\n\nCollaborated in palm oil leaf data collection activities with Koperasi Borneo Agrosindo Sentosa to support dataset development and field validation processes.	/uploads/projects/7c0ad242-5748-4bf9-9e74-dd5a901a1d0a-1780074358.png	https://github.com/ShahDelphi/yolov8-palm-oil-disease-detection.git		Computer Vision, Deep Learning, Python	2	2026-05-30 00:06:00.128673+07	2026-05-30 00:06:00.128673+07	\N
3e0a5cf2-e3b4-42ea-a39b-d79c17e112b2	Deep Learning-Based Sentiment Analysis \nfor Social Media Text Classification	Developed a deep learning-based sentiment analysis system using LSTM architecture for social media text classification.\n\nImplemented NLP preprocessing, topic extraction, and interactive dashboard visualization while integrating Gemini API for enhanced topic modeling and insight generation.	/uploads/projects/c5f0722a-8379-42bb-99fa-765b352d1632-1780074614.png			Natural Language Processing (NLP), Deep Learning, Python, Kecerdasan Buatan (AI), Ilmu Data	3	2026-05-30 00:10:40.212854+07	2026-05-30 00:34:02.223275+07	\N
\.


--
-- TOC entry 3359 (class 0 OID 21404)
-- Dependencies: 213
-- Data for Name: skills; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.skills (id, name, category, "order", created_at, updated_at, deleted_at) FROM stdin;
fddd4052-a5f6-4687-b94a-6ab425226a78	React	Frontend	1	2026-05-28 23:54:33.614934+07	2026-05-28 23:54:33.614934+07	\N
18ddbce9-efcb-4b8d-b477-af98d8a9db16	Tailwind CSS	Frontend	2	2026-05-28 23:54:33.614934+07	2026-05-28 23:54:33.614934+07	\N
76abdf09-3fbc-474e-bf42-de1a9078463f	Vite	Frontend	3	2026-05-28 23:54:33.614934+07	2026-05-28 23:54:33.614934+07	\N
75d21a6c-38b4-4cfc-b654-731d13b7f6c1	HTML5/CSS3	Frontend	4	2026-05-28 23:54:33.614934+07	2026-05-28 23:54:33.614934+07	\N
c586b291-8e23-45b5-9819-77b23450de1b	Golang (Fiber/Gin)	Backend	1	2026-05-28 23:54:33.614934+07	2026-05-28 23:54:33.614934+07	\N
773f79a2-d95f-4be7-85ee-73745ea6ae94	Node.js (Express)	Backend	2	2026-05-28 23:54:33.614934+07	2026-05-28 23:54:33.614934+07	\N
0301dec8-fcbd-44ef-967d-3964c4aac902	RESTful API	Backend	3	2026-05-28 23:54:33.614934+07	2026-05-28 23:54:33.614934+07	\N
05041dda-d563-4609-80e1-2068a069441a	GORM / Hibernate	Backend	4	2026-05-28 23:54:33.614934+07	2026-05-28 23:54:33.614934+07	\N
47de9e81-b690-498a-b35c-b9777eb7d205	Python (Pandas/NumPy)	Data Science & AI	1	2026-05-28 23:54:33.614934+07	2026-05-28 23:54:33.614934+07	\N
325b3cdb-bec9-446e-ab49-89a5e46ddee7	Machine Learning (XGBoost/Scikit-Learn)	Data Science & AI	2	2026-05-28 23:54:33.614934+07	2026-05-28 23:54:33.614934+07	\N
e043c476-2e4c-426b-a43f-da1593706030	Deep Learning (YOLOv8/PyTorch)	Data Science & AI	3	2026-05-28 23:54:33.614934+07	2026-05-28 23:54:33.614934+07	\N
d49d0392-6271-433d-a965-c5a14a468e7d	Forecasting (ARIMA/SARIMAX)	Data Science & AI	4	2026-05-28 23:54:33.621904+07	2026-05-28 23:54:33.621904+07	\N
90f13416-132f-4da8-b273-86581b49e9b0	PostgreSQL	Databases & Tools	1	2026-05-28 23:54:33.621904+07	2026-05-28 23:54:33.621904+07	\N
1b7f31e8-2d41-459e-a860-354899768ee9	MySQL	Databases & Tools	2	2026-05-28 23:54:33.62291+07	2026-05-28 23:54:33.62291+07	\N
77b4d2d9-2919-4762-8d6d-b5f25503c4cd	Git & GitHub	Databases & Tools	3	2026-05-28 23:54:33.62291+07	2026-05-28 23:54:33.62291+07	\N
aa4411d0-685b-4f29-bcb6-ab41ff12c609	Docker	Databases & Tools	4	2026-05-28 23:54:33.62291+07	2026-05-28 23:54:33.62291+07	\N
1c2952a0-c070-4676-ab7e-3413e80a9c86	TensorFlow	Data Science & AI	5	2026-05-29 00:07:02.338295+07	2026-05-29 00:07:02.338295+07	2026-05-29 23:14:51.246352+07
\.


--
-- TOC entry 3195 (class 2606 OID 21367)
-- Name: admins admins_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_pkey PRIMARY KEY (id);


--
-- TOC entry 3208 (class 2606 OID 21402)
-- Name: certificates certificates_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.certificates
    ADD CONSTRAINT certificates_pkey PRIMARY KEY (id);


--
-- TOC entry 3205 (class 2606 OID 21392)
-- Name: experiences experiences_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.experiences
    ADD CONSTRAINT experiences_pkey PRIMARY KEY (id);


--
-- TOC entry 3215 (class 2606 OID 21421)
-- Name: profiles profiles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_pkey PRIMARY KEY (id);


--
-- TOC entry 3203 (class 2606 OID 21381)
-- Name: projects projects_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);


--
-- TOC entry 3212 (class 2606 OID 21412)
-- Name: skills skills_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.skills
    ADD CONSTRAINT skills_pkey PRIMARY KEY (id);


--
-- TOC entry 3198 (class 2606 OID 21369)
-- Name: admins uni_admins_email; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT uni_admins_email UNIQUE (email);


--
-- TOC entry 3200 (class 2606 OID 21371)
-- Name: admins uni_admins_username; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT uni_admins_username UNIQUE (username);


--
-- TOC entry 3196 (class 1259 OID 21372)
-- Name: idx_admins_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_admins_deleted_at ON public.admins USING btree (deleted_at);


--
-- TOC entry 3209 (class 1259 OID 21403)
-- Name: idx_certificates_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_certificates_deleted_at ON public.certificates USING btree (deleted_at);


--
-- TOC entry 3206 (class 1259 OID 21393)
-- Name: idx_experiences_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_experiences_deleted_at ON public.experiences USING btree (deleted_at);


--
-- TOC entry 3213 (class 1259 OID 21422)
-- Name: idx_profiles_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_profiles_deleted_at ON public.profiles USING btree (deleted_at);


--
-- TOC entry 3201 (class 1259 OID 21382)
-- Name: idx_projects_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_projects_deleted_at ON public.projects USING btree (deleted_at);


--
-- TOC entry 3210 (class 1259 OID 21413)
-- Name: idx_skills_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_skills_deleted_at ON public.skills USING btree (deleted_at);


-- Completed on 2026-05-30 00:53:51

--
-- PostgreSQL database dump complete
--

