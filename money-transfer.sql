
CREATE TABLE transfers (
    id SERIAL PRIMARY KEY,                     
    trx_id VARCHAR(50) DEFAULT NULL,
    payment_ref VARCHAR(50) UNIQUE NOT NULL,
    from_account_number VARCHAR(50) unique NOT NULL,          
    to_account_number VARCHAR(50) unique NOT NULL,             
    amount DECIMAL(15, 2) NOT NULL,            
    status VARCHAR(20) NOT NULL,               
    created_at TIMESTAMPTZ DEFAULT NOW(),      
    updated_at TIMESTAMPTZ DEFAULT NOW(),
);