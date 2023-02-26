package serviceImplement

import "bitsys.sys/construction_management_system/bean"

	//if Commit ==1 then and then only data will be save on database after save data in database user cannot be modify it only Contractor can be deleted this record 
	//and only supervisor has authority to add data and modify before save	its applicable to all operations
	//	Connection con=ConnectionProvider.provideConnection();
	c int;
//		DatabaseImplements d=new DatabaseImplements();
	func DataEntry(sitename string,date string,no_of_labor []int,no_of_worker []int,no_of_watchman int,commit int) string {
		var l Labor 
		var w Worker
		var w1 Watchman
		var labor_amt int				//total amount of labor
		var worker_amt int			//total amount of workers
		var watchman_amt int			//total amount of watchman
		var total_amt int				//total amount
		var lcount int				//count total number of labors
		var wcount int				//count total number of worker
		 
		for i:=0; i<len(no_of_labor); i++ {
			lcount+=no_of_labor[i]
		}
		
		for i=0; i<len(no_of_worker); i++ {
			wcount+=no_of_worker[i];
		}
		labor_amt=lcount* ServiceImpleamets.GetRates("labor")			//get rate from rate table and multiply by there counts
		worker_amt=wcount* ServiceImpleamets.GetRates("worker")
		watchman_amt=no_of_watchman* ServiceImpleamets.GetRates("watchman")
		total_amt=labor_amt+worker_amt+watchman_amt
		l= bean.CreateLabor(lcount,labor_amt);
		w= bean.CreateWorker(wcount,worker_amt)
		w1= bean.CreateWatchman(no_of_watchman,watchman_amt);
		if commit == 1 {									//if user wish to sava data in database
		DataEntry(sitename, date,lcount,labor_amt, wcount,worker_amt,no_of_watchman,watchman_amt,total_amt);
		}
		return l+"\n"+w+"\n"+w1+"\ntotal amt="+total_amt;			//print all details for check

		
	}


	func RetriveData(sitename string,commit int) int {
		if commit == 1 {
		c = RetriveData(sitename) }
		return c
	}

	
	func RetriveData(sitename string,date string,commit int) int{
		if commit == 1 {c = RetriveData(sitename, date)}
		
		return c
	}

	
	func DeleteData(sitename string,commit int)	int{
		if commit == 1 {c =  DeleteData(sitename);}
		
		return c
	}

	
	func DeleteData( sitename string,date string,commit int) int {
		if commit == 1 {c = DeleteData(sitename, date)}
		
		return c
	}

	
	func LaborRateChange(rate int,commit int) int {
		if commit == 1 {c = LaborRateChange(rate)}
		
		return c
	}

	
	func WorkerRateChange(int rate,int commit)					//Change in worker rates
	{
		if commit == 1 {c = workerRateChange(rate)}
		return c
	}

	public int cementRateChange(int rate,int commit)					//Change in cement rates 
	{
		if(commit==1)
		c := d.cementRateChange(rate);
		return c;
	}

	public int sandRateChange(int rate,int commit)						//Change in sand rates 
	{
		if(commit==1)
		c := d.sandRateChange(rate);
		return c;
	}


	public int brickRateChange(int rate,int commit)						//Change in Brick rate Changes 
	{
		if(commit==1)
		c := d.brickRateChange(rate);
		return c;
	}


	public String insertUser(String sitename, String role,int commit)		//Insert New User 
	{
		if(commit==1)
		return d.insertUser(sitename, role);
		
		return null;
	}


	public int supervisorSiteChange(int username, String newsite,int commit)		//Supervisor site Change 
	{
		if(commit==1)
		c=d.supervisorSiteChange(username, newsite);
		return c;
	}


	public int customerSiteChange(int username, String password, String sitename,int commit)	// Customer Site change in this case customer password is must
	{
		if(commit==1)
		c=d.customerSiteChange(username, password, sitename);
		return c;
	}
	
	public int watchmanRateChange(int rate,int commit)							//Change in Watchman Rate Change 
	{
		if(commit==1)
			c=d.watchmanRateChange(rate);
		return c;
	}
	
	public int deleteUser(int userid,int commit)								//Delete User 
	{if(commit==1)
		c=d.deleteUser(userid);
	return c;
	}
	
	
	public String retriveUser(int commit)										//Retrieve all User Data
	{
		if(commit==1)
			d.retriveUser();
		return null;
	}
	
	
	public String retriveUser(int userid,int commit)							//Retrieve User data using userid
	{
		if(commit==1)
			d.retriveUser(userid);
		return null;
	}
	
	public String retriveUser(String role,int commit)							//Retrieve User data via there role 
	{
		if(commit==1)
			d.retriveUser(role);
		return null;
	}
	MaterialCal mc=new MaterialCal();
	public String materialEntry(String sitename,String date,int cement_count,int sand_count,int brick_count,int commit)
	{int cement_amt=0;
	int sand_amt=0;
	int brick_amt=0;
	 int total_amt=0;
		try {
		 cement_amt=mc.cementCal(cement_count);
		sand_amt=mc.sandCal(sand_count);
		  brick_amt=mc.brickCal(brick_count);
		  total_amt=cement_amt+sand_amt+brick_amt;
		 if(commit==1)
			 d.materialEntry(sitename, date, cement_count, cement_amt, sand_count, sand_amt, brick_count, brick_amt, total_amt);
	
	} catch (LessThanZeroException e) {
		// TODO Auto-generated catch block
		e.printStackTrace();
	}
	 
	return "cement count= "+cement_count+" cement amount="+cement_amt+" \nsand count"+sand_amt+" sand amount"+sand_amt+"\nbrick count="+brick_count+" brick amount="+brick_amt+"\nTotal amount="+total_amt+"";
	 

		
	}
	
}
