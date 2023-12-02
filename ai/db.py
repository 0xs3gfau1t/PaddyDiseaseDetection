import psycopg
import os
from dotenv import load_dotenv
load_dotenv()

class DbClient:
    def __init__(self):
        self.conn = lambda: psycopg.connect("host={dbhost} port={dbport} dbname={dbname} user={dbuser} password={dbpass}".format(
                        dbname=os.getenv("DB_NAME"),
                        dbuser=os.getenv("DB_USER"),
                        dbpass=os.getenv("DB_PASS"),
                        dbport=os.getenv("DB_PORT"),
                        dbhost=os.getenv("DB_HOST"),
                        autocommit=True)
                    )

    def getImageIdentifier(self, diseaseId: str):
        with self.conn().cursor() as cur:
                cur.execute("SELECT identifier FROM images WHERE image_disease_identified=%s LIMIT 1", (diseaseId,))
                return cur.fetchone()

    def updateDone(self, identifiedId: str, diseaseId: str):
        with self.conn() as conn:
            cur = conn.cursor()
            with conn.transaction():
                cur.execute('UPDATE disease_identifieds SET status=\'processed\' WHERE id=%s', (diseaseId,))
                cur.execute('UPDATE disease_identified_disease SET disease_identified_id=%s, disease_id=%s', (identifiedId, diseaseId))

    def getDiseaseFromName(self, name):
        with self.conn().cursor() as cur:
            cur.execute('SELECT id from diseases where name=%s', (name,))
            return cur.fetchone()

    def test(self, diseaseId):
        print(self.getImageIdentifier(diseaseId))

if __name__ == "__main__":
    DbClient().test("8210287b-1e2a-418d-8f89-433f99d72a2d")
