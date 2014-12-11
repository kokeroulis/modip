package main

import (

	"database/sql"
	"github.com/kokeroulis/modip/models"
)

var departments []*models.Department
var MySQLDb *sql.DB

var L1  map[string]string
var L2  map[string]string
var L3  map[string]string
var L4  map[string]string
var L5  map[string]string
var L6  map[string]string
var L7  map[string]string
var L8  map[string]string
var L9  map[string]string
var L10 map[string]string
var L11 map[string]string
var L12 map[string]string
var L13 map[string]string
var L14 map[string]string
var L15 map[string]string
var L16 map[string]string

var T1  map[string]string
var T2  map[string]string
var T3  map[string]string
var T4  map[string]string
var T5  map[string]string

func init() {
	L1 = map[string]string{}
	L2 = map[string]string{}
	L3 = map[string]string{}
	L4 = map[string]string{}
	L5 = map[string]string{}
	L6 = map[string]string{}
	L7 = map[string]string{}
	L8 = map[string]string{}
	L9 = map[string]string{}
	L10 = map[string]string{}
	L11 = map[string]string{}
	L12 = map[string]string{}
	L13 = map[string]string{}
	L14 = map[string]string{}
	L15 = map[string]string{}
	L16 = map[string]string{}

	T1 = map[string]string{}
	T2 = map[string]string{}
	T3 = map[string]string{}
	T4 = map[string]string{}
	T5 = map[string]string{}

//	L1["Field1"] =
	//L1["Field2"] =

	L2["Field1"] = "field_dep_pnk1_typos"
	L2["Field2"] = "field_dep_pnk1_wres_th"
	L2["Field3"] = "field_dep_pnk1_wres_ap"
	L2["Field4"] = "field_dep_pnk1_wres_e"
	L2["Field5"] = "field_dep_pnk1_ects_th_ap"
	L2["Field6"] = "field_dep_pnk1_ects_e"
	L2["Field7"] = "field_dep_pnk1_oresdid"
	L2["Field8"] = "field_dep_pnk1_oresergnoyes"
	L2["Field9"] = "field_dep_pnk1_oreserg"
	L2["Field10"] = "field_dep_pnk1_pr_wr_dial"
	L2["Field11"] = "field_dep_pnk1_pr_wr_erg"
	L2["Field12"] = "field_dep_pnk1_pr_wr_m_o"
	L2["Field13"] = "field_dep_pnk1_pr_wr_alli"
	L2["Field14"] = "field_dep_pnk1_vivliograyesno"
	L2["Field15"] = "field_dep_pnk1_erg_ypoxr"
	L2["Field16"] = "field_dep_pnk1_erg_proair"

	L3["Field1"] = "field_dep_pnk1_selidaodigsp"
	L3["Field2"] = "field_dep_pnk1_website"
	L3["Field3"] = "field_dep_pnk1_axiol_th"
	L3["Field4"] = "field_dep_pnk1_axiol_e"
	L3["Field5"] = "field_dep_pnk1_axiol"

	L4["Field1"] = "field_dep_pnk1_anapr_ylhs"
	L4["Field2"] = "field_dep_pnk1_epikal_ylhs"

	L5["Field1"] = "field_dep_pnk1_siggrama"
	L5["Field2"] = "field_dep_pnk1_boith"
	L5["Field3"] = "field_dep_pnk1_epikair_boith"
	L5["Field4"] = "field_dep_pnk1_posost_ylhs_boith"
	L5["Field5"] = "field_dep_pnk1_prosth_bibliogr"
	L5["Field6"] = "field_dep_pnk1_gnwst_ylhs"

	L6["Field1"] = "field_dep_pnk1_anak_wr_gr"
	L6["Field2"] = "field_dep_pnk1_methodeysh_ekp"
	L6["Field3"] = "field_dep_pnk1_ekp_episk"

	L7["Field1"] ="field_dep_pnk1_alles_ekp_drast"

	L8["Field1"] = "field_dep_pnk1_symet_foit"
	L8["Field2"] = "field_dep_pnk1_simetoxexet"
	L8["Field3"] = "field_dep_pnk1_simetoxepanex"


	L16["Field1"] = "field_dep_pnk1_sxolio1"

	L15["Field1"] =  "field_dep_pnk1_diadik_aks"
	L15["Field2"] = "field_dep_pnk1_aksiop_erwt_aks"

	L14["Field1"] =  "field_dep_pnk1_kat_bath_04"
	L14["Field2"] =  "field_dep_pnk1_kat_bath_45"
	L14["Field3"] =  "field_dep_pnk1_kat_bath_56"
	L14["Field4"] =  "field_dep_pnk1_kat_bath_67"
	L14["Field5"] =  "field_dep_pnk1_kat_bath_784"
	L14["Field6"] =  "field_dep_pnk1_kat_bath_8510"
	L14["Field7"] =  "field_dep_pnk1_kat_bath_mo"

	L13["Field1"] = "field_dep_pnk1_koin_kat_foit"

	L12["Field1"] = "field_dep_pnk1_ekpmesa"
	L12["Field2"] = "field_dep_pnk1_ekpmesaeparkia"
	L12["Field3"] = "field_dep_pnk1_ekpmesaelipsis"

	L11["Field1"] = "field_dep_pnk1_texnol_plirof"
	L11["Field2"] = "field_dep_pnk1_xrhsh_math_boith"
	L11["Field3"] = "field_dep_pnk1_xr_tpe_erg"
	L11["Field4"] = "field_dep_pnk1_xr_tpe_aks"
	L11["Field5"] = "field_dep_pnk1_xr_tpe_epik"

	L10["Field1"] = "field_dep_pnk1_aith_didask"
	L10["Field2"] = "field_dep_pnk1_erg_math"
	L10["Field3"] = "field_dep_pnk1_diath_erg_ekt_wr"
	L10["Field4"] = "field_dep_pnk1_spoudastiria"
	L10["Field5"] = "field_dep_pnk1_ekp_logismiko"
	L10["Field6"] = "field_dep_pnk1_ikanop_ypost"
	L10["Field7"] = "field_dep_pnk1_kris_diath_ypd"


	L9["Field1"] = "field_dep_pnk1_tropoi_aks"
	L9["Field2"] = "field_dep_pnk1_tropoi_aks_alla"
	L9["Field3"] = "field_dep_pnk1_parak_foit_erg"
	L9["Field4"] = "field_dep_pnk1_lhpsh_sxol_foit"
	L9["Field5"] = "field_dep_pnk1_eksasf_diafan"

	T1["Field1"] = "field_dep_apdid_etos_dimos"
	T1["Field2"] = "field_dep_apdid_biblia_monogr"
	T1["Field3"] = "field_dep_apdid_epist_per_krites"
	T1["Field4"] = "field_dep_apdid_epist_xr_krites"
	T1["Field5"] = "field_dep_apdid_prakt_syn_kr"
	T1["Field6"] = "field_dep_apdid_prakt_syn_xorkr"
	T1["Field7"] = "field_dep_apdid_ana_syn_kr_xr_pr"
	T1["Field8"] = "field_dep_apdid_ana_xr_kr_xr_pr"
	T1["Field9"] = "field_dep_apdid_alles_erg"
	T1["Field10"] = "field_dep_apdid_kef_syl_tom"
	T1["Field11"] = "field_dep_apdid_alla"
	T1["Field12"] = "field_dep_apdid_epeks_alla"
	T1["Field13"] = "field_dep_apdid_epist_dhm"

	T2["Field1"] = "field_dep_apdid_eteroan"
	T2["Field2"] = "field_dep_apdid_anaf_eid_typ"
	T2["Field3"] = "field_dep_apdid_bibliokrisies"
	T2["Field4"] = "field_dep_apdid_sym_ep_epis_syn"
	T2["Field5"] = "field_dep_apdid_sym_syn_epitr"
	T2["Field6"] = "field_dep_apdid_kritis"
	T2["Field7"] = "field_dep_apdid_proskl_die_syn"
	T2["Field8"] = "field_dep_apdid_dipl_eyres"
	T2["Field9"] = "field_dep_apdid_brabeia"
	T2["Field10"] = "field_dep_apdid_timit_titloi"
	T2["Field11"] = "field_dep_apdid_epeks_anagn_erg"

	T3["Field1"] = "field_dep_apdid_ereyn_progr_erg"
	T3["Field2"] = "field_dep_apdid_er_prg_erg_sym"
	T3["Field3"] = "field_dep_apdid_sym_eks_syn_erga"
	T3["Field4"] = "field_dep_apdid_sym_foitht_erga"

	T4["Field1"] = "field_dep_apdid_arith_ergast"
	T4["Field2"] = "field_dep_apdid_epark_erg"
	T4["Field3"] = "field_dep_apdid_epark_erg_eksop"
	T4["Field4"] = "field_dep_apdid_kal_diath_ypod"
	T4["Field5"] = "field_dep_apdid_er_ant_den_kal"
	T4["Field6"] = "field_dep_apdid_entatikh_xr_ypod"
	T4["Field7"] = "field_dep_apdid_syxnh_ananewsh"
	T4["Field8"] = "field_dep_apdid_epid_xrhm_prom"
	T4["Field9"] = "field_dep_apdid_ereyn_synerg"
	T4["Field10"] = "field_dep_apdid_prakt_aksiop"

	T5["Field1"] = "field_dep_apdid_synd_koinwnia"

}
