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