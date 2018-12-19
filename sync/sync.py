import json
import time
import binascii
import pymysql
from ontology.smart_contract.neo_contract.abi.abi_info import AbiInfo
from ontology.utils import util
from ontology.ont_sdk import OntologySdk
from ontology.common.address import Address

rpc_address = 'http://polaris3.ont.io:20336'
sdk = OntologySdk()

def Syncwd():
    sdk.set_rpc(rpc_address)
    blockheight = sdk.rpc.get_block_count()
    lastheight = blockheight


    while 1:
        while(lastheight > blockheight):
            blockheight = sdk.rpc.get_block_count()
            #time.sleep(5)

        while(lastheight <= blockheight):
            print(lastheight)
            #lastheight = 562237
            event = sdk.rpc.get_smart_contract_event_by_height(lastheight-1)
            lastheight += 1
            print(event)
            if event == None:
                continue
            for item in event:
                for state in item['Notify']:
                    if state['ContractAddress'] == '55c494d57ac7cbaad565b5e19ea757dcfe315e29':
                        print(state)
                        if(state['States'][0] != '6775657373'):
                            continue
                        tokentype = state['States'][1]
                        address = state['States'][2]
                        amount = state['States'][3]
                        gameid = state['States'][4]
                        usernumber = state['States'][5]
                        sysnumber = state['States'][6]

                        tokentype = int(tokentype,16)
                        address = Address((binascii.a2b_hex('64f75b59554a2008bcc1a87a7ae09249abc74a91'))).b58encode()
                        amount = util.parse_neo_vm_contract_return_type_integer(amount)
                        usernumber = int(usernumber,16)
                        sysnumber = int(sysnumber,16)
                        gameid = int(gameid,16)
                        print(state['States'])
                        SaveResult(address,tokentype,amount,gameid,usernumber,sysnumber)



def SaveResult(address,tokentype,amount,gameid,usernumber,sysnumber):
    cur = db.cursor()
    try:
        sql_str = ("INSERT INTO guess (address,tokentype,amount,gameid,usernumber,sysnumber ) VALUES ('%s', %s,%s,%s,%s,%s)" % (address, tokentype,amount,gameid,usernumber,sysnumber))
        cur.execute(sql_str)
        db.commit()
    except Exception as e:
        db.rollback()
        raise e
    finally:
        cur.close()
    pass

def connect_wxremit_db():
    return pymysql.connect(host="localhost", user="root",
                         password="123456", db="ontbet", port=3306)


db = connect_wxremit_db()

Syncwd()