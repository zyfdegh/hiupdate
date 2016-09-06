// 8 -> 08
function addZero(str){
    str ='00'+str;
    return str.substring(str.length-2,str.length);
}

// copy
function copy(id){
	var e=document.getElementById(id);
	e.select();
	document.execCommand("Copy");
}

// get yesterday date in 20160906
function getYesterday(){
	return getDate(-1);
}

// if today is Monday, getLastFriday() will return last friday date in 20160906
function getLastFriday(){
	return getDate(-3)
}

// get date in 20160906, days is a interger number of offset
function getDate(days){
	var dd = new Date();
	dd.setDate(dd.getDate()+days);
	var y = dd.getFullYear();
	var m = dd.getMonth()+1;
	var d = dd.getDate();
	return y + addZero(m) + addZero(d);
}

// check if today is Monday
function isMonday(){
	var n = new Date().getDay();
	return n == 1;
}
