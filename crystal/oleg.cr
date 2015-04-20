module Oleg
    @[Link("oleg-http")]
    lib LibOlegHttp
        struct DBConn
            # TODO: How do I pull these from the lib?
            host : UInt8[256]
            port : UInt8[64]
            db_name : UInt8[64]
        end

        fun fetch_data_from_db(conn : DBConn*, key : UInt8[250], outside : UInt64) : UInt8*
    end

    class OlegConnection
        def initialize(host="127.0.0.1", port="38080")
            @conn = LibOlegHttp::DBConn.new
            @conn.host = host[0,256]
            @conn.port = port
            @conn.db_name = "waifu"
        end

        def get(key)
            size :: UInt64
            val = LibOlegHttp.fetch_data_from_db(@conn, key, out size)
        end

        def set(key, value, timeout=nil)
            # return self.connection.set(key, value, timeout)
        end

        def spoil(key, timeout)
        end

        def scoop(key)
            # return self.connection.delete(key)
        end

        def exists(key)
            # return self.connection.has_key(key)
        end

        def bulk_unjar(data)
        end
    end
end
