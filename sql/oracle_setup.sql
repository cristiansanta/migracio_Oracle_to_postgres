-- Crear tabla empleados
CREATE TABLE empleados (
    id NUMBER PRIMARY KEY,
    nombre VARCHAR2(100),
    edad NUMBER,
    salario NUMBER(10,2)
);

-- Insertar datos
INSERT INTO empleados (id, nombre, edad, salario) VALUES (1, 'Juan Pérez', 30, 50000);
INSERT INTO empleados (id, nombre, edad, salario) VALUES (2, 'María García', 35, 60000);

COMMIT;