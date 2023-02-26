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

	
	func WorkerRateChange(rate int,commit int)	int {
		if commit == 1 {c = WorkerRateChange(rate)}
		return c
	}

	func CementRateChange(rate int,commit int)	int {
		if commit==1 {c := CementRateChange(rate)}
		return c
	}

	func SandRateChange(rate int,commit int)	int {
		if commit == 1 {c := SandRateChange(rate)}
		return c
	}

    func BrickRateChange(rate int,commit int)	int {
		if commit == 1 {c := BrickRateChange(rate)}
		return c
	}


	func InsertUser(sitename string,role string,commit int)	int {
		if commit == 1 { return InsertUser(sitename, role) }
		return 0;
	}


	func SupervisorSiteChange(username int,newsite string,commit int)	int {
		if commit == 1 {c = SupervisorSiteChange(username, newsite)}
		return c
	}


	func CustomerSiteChange(username int,password string,sitename string,commit int)	int {
		if commit == 1 {c = CustomerSiteChange(username, password, sitename)}
		return c
	}
	
	func WatchmanRateChange(rate int,commit int)	int {
		if commit == 1 {c = WatchmanRateChange(rate)}
		return c
	}
	
	func DeleteUser(userid int,commit int) int  {
		if commit == 1 {c = DeleteUser(userid)}
	return c
	}
	
	
	func RetriveUser(commit int) string {
		if commit == 1 {return RetriveUser()}
		return nil
	}
	func RetriveUser(userid int,commit int)	string {
		if commit == 1 {return RetriveUser(userid)}
		return nil
	}
	
	func RetriveUser(role string,commit int) string {
		if commit == 1 {return RetriveUser(role)}
		return nil;
	}
	var mc MaterialCal = CreateMaterialCal();
	func MaterialEntry(String sitename,String date,int cement_count,int sand_count,int brick_count,int commit) string {
	var cement_amt int
	var sand_amt int
	var brick_amt int
	var total_amt int
		try {
		 cement_amt= mc.cementCal(cement_count);
		sand_amt= mc.sandCal(sand_count);
		  brick_amt = mc.brickCal(brick_count);
		  total_amt = cement_amt+sand_amt+brick_amt;
		 if commit == 1 {
			 d.materialEntry(sitename, date, cement_count, cement_amt, sand_count, sand_amt, brick_count, brick_amt, total_amt);
		}
	} catch (LessThanZeroException e) {
		// TODO Auto-generated catch block
		e.printStackTrace();
	}
	 
	return "cement count= "+cement_count+" cement amount="+cement_amt+" \nsand count"+sand_amt+" sand amount"+sand_amt+"\nbrick count="+brick_count+" brick amount="+brick_amt+"\nTotal amount="+total_amt+"";
	 

		
	}
	
}
