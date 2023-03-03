package aspect

type ConnectionProvider struct {
	//static Connection con;
	//static FileInputStream fis;
	//static String s[]=new String[4];
	//static Properties p;
}

func GetProperties() {
	/*	try {
			fis=new FileInputStream(".//resources//DBConfig.properties");
			int x=0;
			p=new Properties();
			p.load(fis);


		s[0]=p.getProperty("DriverName");
		s[1]=p.getProperty("Url");
		s[2]=p.getProperty("UserName");
		s[3]=p.getProperty("Password");
			}catch(Exception e) {
				e.getStackTrace();
			} */
}

func ProvideConnection() {
	/*	try {
			getProperties();
			Class.forName(s[0]);
			con=DriverManager.getConnection(s[1],s[2],s[3]);
			System.out.println("Connection Established");
		} catch (ClassNotFoundException e) {
			e.printStackTrace();
		} catch (SQLException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
	return con; */

}
