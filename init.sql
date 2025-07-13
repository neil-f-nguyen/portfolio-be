-- Drop existing tables if they exist
DROP TABLE IF EXISTS fun_facts CASCADE;
DROP TABLE IF EXISTS goals CASCADE;
DROP TABLE IF EXISTS learning_items CASCADE;
DROP TABLE IF EXISTS about_cards CASCADE;
DROP TABLE IF EXISTS experiences CASCADE;
DROP TABLE IF EXISTS skills CASCADE;
DROP TABLE IF EXISTS projects CASCADE;
DROP TABLE IF EXISTS profiles CASCADE;

-- Create database schema
CREATE TABLE profiles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(50),
    location VARCHAR(255),
    website VARCHAR(255),
    github VARCHAR(255),
    linkedin VARCHAR(255),
    avatar VARCHAR(500),
    about TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    image TEXT[],
    technologies TEXT[],
    github VARCHAR(255),
    demo VARCHAR(255),
    featured BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE skills (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    level VARCHAR(50) NOT NULL,
    icon_url VARCHAR(500),
    category VARCHAR(100) NOT NULL
);

CREATE TABLE experiences (
    id SERIAL PRIMARY KEY,
    company VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    location VARCHAR(255),
    start_date DATE NOT NULL,
    end_date DATE,
    current BOOLEAN DEFAULT false,
    description TEXT,
    technologies TEXT[],
    logo VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- New tables for About section
CREATE TABLE about_cards (
    id SERIAL PRIMARY KEY,
    icon VARCHAR(10) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    gradient_colors TEXT[] NOT NULL,
    display_order INTEGER NOT NULL,
    hover_rotate_y INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE learning_items (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    color VARCHAR(50) NOT NULL,
    display_order INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE goals (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    color VARCHAR(50) NOT NULL,
    display_order INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE fun_facts (
    id SERIAL PRIMARY KEY,
    emoji VARCHAR(10) NOT NULL,
    text VARCHAR(255) NOT NULL,
    rotate_angle INTEGER DEFAULT 0,
    display_order INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert sample data
INSERT INTO profiles (name, title, email, phone, location, website, github, linkedin, avatar, about) VALUES
('Nguyen Anh Tu', 'Frontend Developer', 'tu.nguyen@example.com', '+84 123 456 789', 'Hanoi, Vietnam', 'https://tuportfolio.com', 'https://github.com/tunguyen', 'https://linkedin.com/in/tunguyen', 'https://dummyimage.com/200x200/000/fff&text=TU', 'Frontend Developer with 3 years of experience in web development. Specialized in React, TypeScript and Next.js.');

INSERT INTO projects (title, description, image, technologies, github, demo, featured) VALUES
('E-Commerce Platform', 'Modern e-commerce platform with online payment and order management.', ARRAY['https://dummyimage.com/600x400/000/fff&text=E-Commerce'], ARRAY['React', 'TypeScript', 'Node.js', 'PostgreSQL', 'Stripe'], 'https://github.com/tunguyen/ecommerce', 'https://ecommerce.tuportfolio.com', true),
('Task Management App', 'Task management application with intuitive interface and collaboration features.', ARRAY['https://dummyimage.com/600x400/000/fff&text=Task+App'], ARRAY['Next.js', 'TypeScript', 'TailwindCSS', 'Prisma', 'PostgreSQL'], 'https://github.com/tunguyen/taskapp', 'https://taskapp.tuportfolio.com', false),
('Portfolio Website', 'Personal portfolio website with responsive design and modern UI.', ARRAY['https://dummyimage.com/600x400/000/fff&text=Portfolio'], ARRAY['Next.js', 'TypeScript', 'TailwindCSS', 'Framer Motion'], 'https://github.com/tunguyen/portfolio', 'https://tuportfolio.com', true);

INSERT INTO skills (name, level, icon_url, category) VALUES
-- Frontend
('React', 'Advanced', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/react/react-original.svg', 'Frontend'),
('Next.js', 'Advanced', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/nextjs/nextjs-original.svg', 'Frontend'),
('TypeScript', 'Advanced', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/typescript/typescript-original.svg', 'Frontend'),
('JavaScript', 'Advanced', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/javascript/javascript-original.svg', 'Frontend'),
('HTML5', 'Advanced', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/html5/html5-original.svg', 'Frontend'),
('CSS3', 'Advanced', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/css3/css3-original.svg', 'Frontend'),
('TailwindCSS', 'Advanced', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/tailwindcss/tailwindcss-plain.svg', 'Frontend'),
('Sass', 'Intermediate', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/sass/sass-original.svg', 'Frontend'),
-- Backend
('Node.js', 'Intermediate', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/nodejs/nodejs-original.svg', 'Backend'),
('Express.js', 'Intermediate', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/express/express-original.svg', 'Backend'),
('PostgreSQL', 'Intermediate', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/postgresql/postgresql-original.svg', 'Backend'),
('MongoDB', 'Intermediate', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/mongodb/mongodb-original.svg', 'Backend'),
-- Tools
('Git', 'Advanced', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/git/git-original.svg', 'Tools'),
('Docker', 'Intermediate', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg', 'Tools'),
('VS Code', 'Advanced', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/vscode/vscode-original.svg', 'Tools'),
('Figma', 'Intermediate', 'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/figma/figma-original.svg', 'Tools');

INSERT INTO experiences (company, position, location, start_date, end_date, current, description, technologies, logo) VALUES
('TechCorp Solutions', 'Frontend Developer', 'Hanoi, Vietnam', '2022-01-01', NULL, true, 'Developed and optimized UI for web products. Worked with React, TypeScript and Next.js.', ARRAY['React', 'Next.js', 'TypeScript', 'TailwindCSS', 'Redux'], 'https://dummyimage.com/100x100/000/fff&text=TC'),
('Digital Innovations', 'Junior Frontend Developer', 'Hanoi, Vietnam', '2021-03-01', '2021-12-31', false, 'Developed user interfaces for web applications. Learned and applied new technologies.', ARRAY['React', 'JavaScript', 'CSS3', 'HTML5', 'Bootstrap'], 'https://dummyimage.com/100x100/000/fff&text=DI'),
('StartupXYZ', 'Intern Frontend Developer', 'Hanoi, Vietnam', '2020-06-01', '2021-02-28', false, 'Frontend intern, supported development of basic features for company website.', ARRAY['HTML5', 'CSS3', 'JavaScript', 'jQuery'], 'https://dummyimage.com/100x100/000/fff&text=SX');

-- Insert About Cards (What Drives Me)
INSERT INTO about_cards (icon, title, description, gradient_colors, display_order, hover_rotate_y) VALUES
('🎨', 'Creative Design', 'Creating beautiful interfaces and excellent user experiences', ARRAY['#3B82F6', '#8B5CF6', '#EC4899'], 1, 10),
('⚡', 'Performance', 'Optimizing performance and page load speed for smooth experiences', ARRAY['#F59E0B', '#EF4444', '#8B5CF6'], 2, -10),
('🚀', 'Innovation', 'Always exploring and applying the latest technologies in web development', ARRAY['#10B981', '#3B82F6', '#8B5CF6'], 3, 8),
('🧠', 'Problem Solving', 'Analyzing and solving complex problems logically and effectively', ARRAY['#6366F1', '#8B5CF6', '#EC4899'], 4, -8),
('🤝', 'Collaboration', 'Working effectively in teams and sharing knowledge with colleagues', ARRAY['#FB7185', '#EC4899', '#8B5CF6'], 5, 6),
('📚', 'Continuous Learning', 'Constantly learning and updating knowledge in the technology field', ARRAY['#34D399', '#14B8A6', '#06B6D4'], 6, -6),
('💡', 'Creative Thinking', 'Creative thinking and providing unique solutions for each project', ARRAY['#8B5CF6', '#6366F1', '#3B82F6'], 7, 5),
('🎯', 'Goal Oriented', 'Always focusing on goals and completing work efficiently', ARRAY['#F59E0B', '#F97316', '#EF4444'], 8, -5);

-- Insert Learning Items
INSERT INTO learning_items (title, color, display_order) VALUES
('Next.js 14 & App Router', 'blue', 1),
('TypeScript Advanced', 'purple', 2),
('Three.js & WebGL', 'pink', 3);

-- Insert Goals
INSERT INTO goals (title, color, display_order) VALUES
('Build scalable web apps', 'green', 1),
('Contribute to open source', 'yellow', 2),
('Share knowledge via blog', 'red', 3);

-- Insert Fun Facts
INSERT INTO fun_facts (emoji, text, rotate_angle, display_order) VALUES
('☕', 'Coffee addict\n3+ cups/day', -2, 1),
('🎵', 'Code with\nlo-fi music', 2, 2),
('🌙', 'Night owl\ndeveloper', -1, 3),
('🎮', 'Gaming for\ninspiration', 1, 4); 