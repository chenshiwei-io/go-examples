package main

import (
	"fmt"
	"strings"
	"unicode"
)

var orderIDs = map[int]int{
	3826430:  0,
	4543794:  0,
	4594990:  0,
	5132573:  0,
	5132462:  0,
	5336999:  0,
	5418303:  0,
	5535055:  0,
	5629207:  0,
	5533218:  0,
	5641281:  0,
	5838739:  0,
	6036673:  0,
	6127471:  0,
	6247812:  0,
	6473342:  0,
	6515971:  0,
	6546011:  0,
	6584129:  0,
	6508598:  0,
	6866192:  0,
	7070591:  0,
	7138330:  0,
	7150072:  0,
	7127502:  0,
	7132035:  0,
	6796848:  0,
	7187278:  0,
	7482033:  0,
	7130002:  0,
	7522572:  0,
	7508213:  0,
	7612477:  0,
	21927886: 0,
	23972162: 0,
	18215341: 0,
	19317640: 0,
	21721267: 0,
	52275620: 0,
	54429305: 0,
	54069046: 0,
	64293211: 0,
	51386673: 0,
	31202241: 0,
	31814792: 0,
	31829519: 0,
	32426766: 0,
	34189007: 0,
	35115945: 0,
	35712937: 0,
	35923048: 0,
	35856446: 0,
	36006864: 0,
	36113123: 0,
	36966814: 0,
	36938831: 0,
	37373888: 0,
	37312847: 0,
	37948841: 0,
	38707975: 0,
	39112495: 0,
	41779514: 0,
	43857281: 0,
	45146389: 0,
	45970813: 0,
	46688455: 0,
	47359768: 0,
	49266227: 0,
	49987852: 0,
	72593646: 0,
	74548821: 0,
	74074096: 0,
	73181092: 0,
	76044313: 0,
	76943765: 0,
	76944914: 0,
	77484841: 0,
	77681565: 0,
	76718994: 0,
	75336355: 0,
	78706540: 0,
	78720547: 0,
	79136146: 0,
	73525167: 0,
	82450584: 0,
	81398908: 0,
	82527935: 0,
	73079745: 0,
	82790325: 0,
	82090468: 0,
	82622239: 0,
	77032537: 0,
	82546121: 0,
	82710890: 0,
	83121650: 0,
	82905602: 0,
	83297611: 0,
	77704595: 0,
	81689037: 0,
	82987903: 0,
	83210800: 0,
	83323668: 0,
	83620050: 0,
	83254739: 0,
	84438878: 0,
	83007380: 0,
	84957117: 0,
	79880965: 0,
	82760786: 0,
	84874770: 0,
	85355332: 0,
	85764112: 0,
	86236780: 0,
	85835896: 0,
	84660987: 0,
	86671647: 0,
	82325350: 0,
	85210233: 0,
	86312567: 0,
	86774306: 0,
	86831001: 0,
	83902372: 0,
	86648000: 0,
	86856214: 0,
	85285757: 0,
	86470639: 0,
	86048837: 0,
	86610937: 0,
	86406015: 0,
	86013573: 0,
	86876808: 0,
	86636754: 0,
	86782798: 0,
	82096621: 0,
	86876441: 0,
	86781020: 0,
	86603233: 0,
	76896086: 0,
	83574285: 0,
	85268802: 0,
	86298342: 0,
	85600555: 0,
	76383945: 0,
	86480070: 0,
	86739876: 0,
	84291712: 0,
	86236698: 0,
	86505486: 0,
	85119670: 0,
	85564462: 0,
	86033338: 0,
	86054383: 0,
	86245781: 0,
	83657068: 0,
	86470861: 0,
	84242935: 0,
	82928575: 0,
	83501718: 0,
	86510194: 0,
	86898644: 0,
	86130637: 0,
	86240599: 0,
	86736702: 0,
	86079073: 0,
	86503836: 0,
	82059461: 0,
	16816422: 0,
	16817186: 0,
	17521855: 0,
	20585912: 0,
	20928387: 0,
	21907354: 0,
	22484737: 0,
	25551620: 0,
	26833638: 0,
}

var s = []int{5710, 5927, 5979, 5987, 6000, 6010, 6012, 6095, 6266, 6334, 6746, 7555, 7581, 7626, 7887, 8153, 8386, 8447, 8598, 8805, 8837, 9371, 9507, 9567, 9890, 9974, 10590, 10676, 10855, 11107, 11238, 11375, 11855, 12201, 12211, 12278, 12885, 13116, 13313, 13462, 13637, 13701, 14450, 14529, 14896, 14901, 14936, 14994, 15192, 15309, 15419, 15646, 15785, 15936, 16125, 16163, 16221, 16245, 16290, 16304, 16376, 16452, 16510, 16686, 16767, 17299, 17354, 17595, 17666, 17714, 17782, 17793, 17802, 18059, 18131, 18208, 18250, 18364, 18490, 18703, 18933, 18969, 19109, 19280, 19305, 19377, 19465, 19467, 19497, 19576, 19681, 19696, 19746, 19837, 20211, 20475, 20577, 20579, 20664, 20923, 21272, 21593, 21670, 21856, 22018, 22070, 22074, 22136, 22222, 22236, 22293, 22340, 22647, 22716, 22904, 23392, 23604, 23728, 23850, 23897, 23908, 24022, 24168, 24178, 24277, 24423, 24473, 24729, 24795, 24842, 24895, 25190, 25333, 25438, 25489, 25854, 25958, 25977, 26168, 26174, 26288, 26306, 26464, 26546, 26680, 26753, 26813, 26996, 27139, 27273, 27416, 27505, 27584, 27918, 28151, 28247, 28293, 28381, 28388, 28391, 28517, 28552, 28554, 28560, 28658, 28660, 28787, 28821, 28849, 28884, 28919, 29061, 29199, 29223, 29266, 29354, 29437, 29772, 29831, 29851, 29983, 30024, 30171, 30217, 30294, 30447, 30569, 30721, 30817, 30825, 30919, 31116, 31484, 31511, 31723, 32064, 32207, 32413, 32473, 32880, 33023, 33056, 33238, 33240, 33351, 33596, 33770, 33830, 33870, 34143, 34396, 34478, 34507, 34540, 34592, 35480, 35510, 35609, 35899, 35966, 35983, 36195, 36275, 36293, 36426, 36437, 36481, 36515, 36574, 36612, 36694, 36944, 37088, 37123, 37202, 37203, 37223, 37271, 37418, 37456, 37468, 37566, 37579, 37590, 37763, 37781, 37794, 37852, 37882, 37912, 37914, 37920, 38011, 38036, 38214, 38258, 38410, 38607, 38640, 38694, 38743, 38821, 39009, 39349, 39440, 39554, 39815, 39850, 40212, 40280, 40444, 40626, 40635, 41106, 41197, 41320, 41388, 41505, 41631,
	41775, 41845, 42559, 42662, 43295, 43450, 43675, 43741, 44011, 44052, 44053, 44201, 44332, 44436, 44446, 44458, 44480, 44631, 44644, 44852, 45083, 45091, 45567, 45875, 46013, 46249, 46693, 46742, 46746, 46849, 46891, 46939, 46989, 47146, 47435, 47480, 47860, 47978, 48001, 48049, 48266, 48418, 48591, 48638, 48677, 49229,
	49238, 49445, 49550, 49618, 49627, 49630, 49655, 49804, 49991, 50047, 50057, 50123, 50184, 50192, 50193, 50333, 50422, 50603, 50901, 50904, 50917, 50931, 50939, 50941, 51086, 51114, 51155, 51170, 51270, 51277, 51302, 51369, 51424, 51455, 51553, 51754, 51788, 51789, 51802, 51821, 51959, 52044, 52053, 52142, 52162, 52163, 52382, 52509, 52534, 52643, 52790, 52880, 52887, 53081, 53115, 53135, 53142, 53253, 53287, 53352, 53357, 53360, 53595, 53610, 53616, 54087, 54192, 54219, 54352, 54414, 54460, 54671, 54730, 54765, 54781, 54864, 54873, 54878, 54903, 54930, 54935, 55148, 55150, 55159, 55184, 55189, 55191, 55198, 55199, 55202,
	55203, 55211, 55221, 55283, 55348, 55359, 55431, 55471, 55565, 55581, 55615, 55690, 55707, 55851, 55909, 55984, 56132, 56141, 56152, 56172, 56312, 56367, 56373, 56384, 56475, 56517, 56570, 56577, 56597, 56609, 56637, 56640, 56646, 56658, 56686, 56726, 56789, 56807, 56812, 56813, 56861, 56904, 57072, 57107, 57370, 57439,
	57490, 57680, 57871, 57924, 57932, 58401, 58484, 58498, 58604, 58696, 58807, 58813, 58905, 59034, 59119, 59225, 59287, 59463, 59503, 59545, 59561, 59798, 59807, 59818, 59833, 59876, 60077, 60154, 60179, 60378, 60403, 60527, 60585, 60980, 61002, 61026, 61040, 61068, 61139, 61179, 61312, 61371, 61398, 61665, 61700,
	61794, 61796, 61828, 61911, 62164, 62233, 62425, 62447, 62452, 62605, 62769, 62980, 62981, 62990, 63029, 63080, 63099, 63194, 63200, 63316, 63398, 63454, 63605, 63928, 64030, 64065, 64075, 64108, 64268, 65007, 65061, 65189, 65216, 65250, 65285, 65471, 65502, 65639, 65720, 66000, 66055, 66067, 66195, 66244, 66307,
	66373, 66379, 66495, 66507, 66531, 66607, 66615, 66626, 66809, 66887, 66911, 67002, 67014, 67057, 67065, 67103, 67159, 67227, 67387, 67500, 67505, 67550, 67645, 67693, 67732, 67853, 67879, 67978, 68041, 68062, 68075, 68118, 68141, 68179, 68190, 68240, 68325, 68483, 68505, 68589, 68620, 68639, 68763, 68775, 68808,
	68910,
	69128, 69141, 69201, 69344, 69345, 69418, 69425, 69492, 69634, 69741, 69860, 69872, 69889, 69957, 70127, 70516, 70529, 70540, 70545, 70669, 70691, 70720, 70753, 70857, 70875, 70960, 70972, 71144, 71203, 71220, 71487, 71516, 71603, 71660, 71725, 71994, 72298, 72552, 72783, 72949, 73315, 73430, 73527, 73529, 73706,
	73875, 74243, 74254, 74399, 74768, 75053, 75560, 75756, 75808, 75842, 75843, 75919, 75987, 75991, 75992, 76052, 76244, 76297, 76468, 76810, 76839, 76891, 76897, 76922, 76939, 77061, 77186, 77191, 77246, 77326, 77476, 77665, 77873, 78014, 78038, 78189, 78259, 78474, 78574, 78598, 78692, 78703, 78801, 78806, 78906,
	78963, 79001, 79014, 79018, 79029, 79043, 79044, 79051, 79165, 79333, 79342, 79352, 79419, 79471, 79771, 80058, 80162, 80319, 80357, 80407, 80412, 80570, 80622, 80849, 81172, 81225, 81287, 81294, 81516, 81760, 81959, 82208, 82233, 82296, 83022, 83054, 83136, 83406, 83738, 83809, 83816, 83845, 83868, 83922, 83946,
	83972, 84008, 84055, 84221, 84311, 84427, 84507, 84713, 84730, 84751, 84809, 85006, 85034, 85121, 85185, 85196, 85206, 85243, 85251, 85299, 85313, 85361, 85376, 85379, 85389, 85505, 85508, 85579, 85635, 85643, 85702, 85716, 85719, 85858, 86033, 86087, 86189, 86379, 86464, 86512, 86780, 86797, 86913, 86969, 86972, 87022,
	87051, 87111, 87196, 87197, 87329, 87360, 87380, 87402, 87409, 87415, 87461, 87511, 87569, 87576, 87580, 87585, 87591, 87602, 87606, 87615, 87625, 87634, 87640, 87647, 87741, 87757, 87873, 88037, 88199, 88220, 88254, 88431, 88443, 88483, 88495, 88506, 88543, 88578, 88589, 88600, 88670, 88719, 88775, 88837, 89099,
	89143, 89237, 89699, 89744, 90053, 90064, 90374, 90407, 90450, 90726, 90787, 90800, 91321, 91322, 91568, 91926, 92342, 92485, 92795, 92942, 93276, 93376, 93599, 94068, 94283, 94286, 94582, 94780, 95203, 95395, 95503, 95572, 95901, 96070, 96279, 96500, 96609, 96633, 96646, 96691, 96727, 96977, 97081, 97200, 97229, 97398,
	97415, 97471, 97684, 97741, 97963, 97999, 98026, 98172, 98279, 98608, 98761, 98928, 99633, 99656, 99800, 100081, 100087, 100171, 100188, 100348, 100586, 100647, 100799, 100817, 100819, 101075, 101500, 101527, 101668, 101681, 101709, 101714, 102038, 102089, 102166, 102348, 102543, 102567, 102590, 102818, 102832,
	102915, 102928, 103038, 103071, 103491, 104060, 104071, 104151, 104174, 104292, 104436, 104698, 104744, 104861, 104926, 104951, 105051, 105130, 105144, 105158, 105218, 105864, 105920, 106121, 106503, 106610, 106685, 106704, 106759, 106802, 106805, 106811, 106824, 106877, 106892, 107327, 107371, 107459, 107482,
	107547, 107572, 107952, 108057, 108197, 108538, 108623, 108875, 109176, 109270, 109286, 109442, 109473, 109491, 109510, 109588, 109782, 109848, 109919, 109984, 110045, 110138, 110233, 110268, 110302, 110303, 110375, 110458, 110462, 110463, 110535, 110606, 110663, 110838, 110851, 110867, 110930, 110932,
	110939, 111052, 111080, 111148, 111169, 111179, 111251, 111288, 111310, 111339, 111341, 111348, 111349, 111351}

func main() {

	s1 := "chen shiwei hello world!"
	for key, value := range strings.Fields(s1) {
		fmt.Printf("k %v v %v\n", key, value)
	}

	for key, value := range strings.FieldsFunc(s1, func(r rune) bool {
		return !unicode.IsLetter(r)
	}) {
		fmt.Printf("k %v v %v\n", key, value)
	}

	fmt.Println(strings.Repeat("a", 3))
	return

	//for i := 0; i < 4; i++ {
	//	s = append(s, s...)
	//}
	//b, _ := json.Marshal(s)
	//
	//fmt.Println(string(b))
	//
	//return

	var a = [5]int{1, 2, 3, 4, 5}

	s := a[0:5]

	if len(s) > 4 {
		s = s[:4]
	}

	fmt.Printf("a: 指针%p cap:%d len:%d %v s:指针%p cap:%d len:%d %v \n", &a, cap(a), len(a), a, s, cap(s), len(s), s)

	s = append(s, 6)

	fmt.Printf("a: 指针%p cap:%d len:%d %v s:指针%p cap:%d len:%d %v \n", &a, cap(a), len(a), a, s, cap(s), len(s), s)

	return

	s = append(s, 7, 8, 9)

	fmt.Printf("a: 指针%p cap:%d len:%d %v s:指针%p cap:%d len:%d %v \n", &a, cap(a), len(a), a, s, cap(s), len(s), s)

}