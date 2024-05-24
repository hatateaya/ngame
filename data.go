package main

type Data struct {
	player_name string
	// basic
	money         int
	energy        int
	proactivition int
	// psy
	selfdistory    int
	gd             int
	dissociation   int
	depressanxiety int
	dyssomnia      int
	nervosa        int
	ocd            int
	phobia         int
	did            int
	amnesia        int
	hallucination  int
	delirium       int
	// od med
	med_dxm  int
	med_bron int
	med_pr80 int
	med_bzd  int
	med_admt int
	med_dph  int
	// status
	dxm_resi   int
	med_use    int
	is_hrt     bool
	libido     int
	transation int
	scar       int
	// ability
	sex_abi             int
	pharma_abi          int
	chem_abi            int
	social_disabled_abi int
	social_abled_abi    int
	//record
	last_med_use int
}
