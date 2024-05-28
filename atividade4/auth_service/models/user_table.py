from sqlalchemy.orm import DeclarativeBase
from sqlalchemy import Column, Integer, String, MetaData, create_engine

metadata = MetaData()
engine = create_engine("mysql+pymysql://admin:admin@database:3306/database")

class Base(DeclarativeBase):
    pass 

class UserTable(Base):
    __tablename__ = "users"
    metadata
    id = Column(Integer,primary_key=True, autoincrement=True)
    email = Column(String(100), nullable=False)
    password = Column(String(100), nullable=False)

metadata.create_all(bind=engine, tables=[UserTable.__table__])
conn = engine.connect()
