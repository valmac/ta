package tests

import (
	"math"
	"testing"

	"../../ta"
)

func TestSTOCHCalculation(t *testing.T) {
	var closes []float64 = []float64{
		1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
		1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
		1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
		1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
		1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
		1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
		1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
		1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
		1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
		1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
	}

	var high []float64 = []float64{
		1.31463198, 2.24011529, 4.35090395, 5.72306104, 5.83616541, 7.48720048, 10.07346761, 10.93418835, 12.82025593, 12.79720793,
		1.42985701, 2.27846961, 3.52112016, 5.79532125, 7.25103722, 8.24901707, 9.38207988, 11.50994503, 10.21403220, 10.51589817,
		1.02628542, 2.80291187, 4.02030305, 4.76219538, 5.14120353, 8.65658136, 7.08021805, 10.95169978, 10.02914505, 13.34340464,
		1.17785031, 2.09943531, 4.30154472, 4.63942146, 6.71578780, 6.46164116, 9.93010304, 9.30014225, 12.86390307, 13.23325131,
		1.43450569, 2.49318175, 4.22576423, 5.90397482, 5.79021058, 6.15732131, 9.24122551, 8.03691993, 11.10328061, 14.94075893,
		1.15158535, 2.82058528, 3.02751880, 4.94961471, 6.21600725, 6.67537735, 9.33445194, 11.42713086, 9.49141009, 14.67260786,
		1.26604057, 2.36399110, 3.54412130, 5.13076388, 7.04124678, 8.21237812, 8.27139865, 8.03620951, 12.58231610, 10.36083825,
		1.16648531, 2.62406172, 4.46506809, 5.84697914, 6.35394185, 6.66025229, 7.72887437, 9.55836813, 12.58454343, 12.47424721,
		1.32325231, 2.25612161, 4.42149232, 4.06963890, 6.45508020, 7.84561638, 9.12005648, 9.86539953, 9.66460971, 10.20838452,
		1.46028771, 2.01339111, 3.63189769, 4.02448126, 5.53118350, 6.22192855, 7.82845249, 9.54736332, 12.07228213, 12.03786138,
	}

	var low []float64 = []float64{
		0.72616039, 1.80597976, 2.20139647, 2.53678620, 4.16885098, 3.44215391, 6.55435985, 7.93489565, 6.76890429, 7.29221287,
		0.91264056, 1.60203779, 2.42122691, 2.92867279, 3.57226757, 5.66056531, 3.53726890, 7.35290217, 5.65582567, 7.58342897,
		0.63110584, 1.14492996, 2.10050288, 3.95820571, 3.22983611, 3.10293001, 5.55146593, 7.91686284, 4.63230507, 7.97904186,
		0.89723676, 1.28874940, 2.85631481, 2.58173746, 4.60274198, 5.77233473, 4.26792994, 4.85523434, 6.24236459, 9.90154062,
		0.98456041, 1.38947772, 1.72424913, 3.31472229, 3.61431918, 4.77256160, 6.67210185, 7.47541841, 5.39501161, 6.40104112,
		0.61003153, 1.42636809, 1.54018488, 3.65000465, 4.61798633, 5.22832228, 3.68839431, 7.58274101, 8.68113834, 7.92514191,
		0.83799467, 1.18254225, 2.25926941, 3.06804425, 2.75277825, 4.48944850, 5.93398517, 6.98592001, 6.49252055, 9.44454695,
		0.64929779, 1.86333510, 2.65703581, 3.01404701, 2.68656815, 3.06354667, 6.24786292, 6.63780861, 8.91809823, 9.72840865,
		0.88275922, 1.29841046, 1.70404561, 2.68502497, 3.45298574, 4.17757103, 3.57606140, 4.48746603, 5.64599136, 5.41777762,
		0.85530878, 1.84998818, 2.93309517, 2.46152791, 2.99014222, 3.22147039, 5.47387887, 4.76199605, 5.32240416, 8.56152385,
	}

	var expected_slowk []float64 = []float64{
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), 27.87890027, 50.85070959, 55.90401438, 60.48119718, 59.87263969, 64.89993016,
		32.81678680, 20.97420866, 19.25191313, 22.61489989, 43.55029983, 54.85756243, 60.31823535, 59.70792505, 64.11304324, 72.58702226,
		37.98897768, 25.28603400, 24.62553560, 29.35356162, 63.11136538, 63.87261130, 69.30237086, 65.84760304, 70.49064424, 68.92086080,
		34.87326125, 21.86676225, 19.38081354, 22.15513997, 46.33345424, 66.57208242, 63.34893405, 68.54159402, 61.79572053, 62.86587247,
		31.49596171, 19.89307521, 18.17370084, 21.39608255, 51.51021174, 74.10528969, 72.14488542, 75.60062262, 73.75777263, 62.58372971,
		32.65246736, 21.17584260, 18.92653665, 21.29089011, 49.79981076, 68.46652608, 69.25782069, 62.59550901, 65.61605476, 61.53837599,
		31.35469526, 19.87697835, 17.75224334, 20.30399929, 43.69896224, 56.11384973, 67.48331110, 81.28271989, 72.41915990, 70.25530869,
		36.59711580, 23.95807411, 21.82860800, 28.16544054, 52.21570600, 69.22580497, 77.42739804, 77.37483486, 70.57995945, 71.71711810,
		36.35951163, 22.95356108, 20.52343847, 23.70796869, 48.79766360, 60.30414467, 65.85829286, 69.93957663, 78.22194298, 87.53999465,
		44.54349272, 28.39107679, 25.66070386, 29.64135244, 59.14063204, 77.03221473, 80.79797682, 79.48028066, 72.82627333, 74.70640179,
	}

	var expected_slowd []float64 = []float64{
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), 9.29296676, 30.07183817, 42.98792627, 51.73456173, 55.80360071, 60.35176543,
		46.58427612, 33.77924239, 26.51557776, 24.56523883, 34.05776933, 44.45766588, 52.38795061, 56.04793783, 60.08049053, 66.33375640,
		52.16136704, 38.72370052, 31.67461806, 30.51408984, 46.81272761, 55.34266945, 62.32252016, 64.08506160, 67.28785292, 68.10435686,
		51.48880905, 36.67778565, 28.02929960, 25.09221978, 35.71283701, 51.14245972, 57.24569689, 62.89364545, 62.34468299, 62.60527773,
		47.05061972, 33.47184746, 25.82277415, 23.60942835, 37.55982004, 55.83255487, 63.98872014, 69.79467138, 71.77622201, 67.17997586,
		49.91622161, 35.54603211, 27.23628438, 24.26358724, 37.03169900, 52.74911254, 61.00346662, 61.79948782, 63.70777129, 62.62307364,
		46.98888445, 33.43293140, 25.59258737, 22.94829333, 33.32362779, 44.71873876, 56.10102493, 68.69187241, 70.55551616, 70.40541242,
		53.50126411, 38.72966911, 30.27913856, 29.22228955, 40.71899777, 54.97240137, 66.19989971, 71.78736728, 71.18366337, 71.45039074,
		53.90495118, 38.42925613, 29.47634730, 26.59215799, 37.69491080, 48.99952773, 57.42891030, 63.68424346, 70.95309322, 79.24654394,
		61.89501833, 45.14304756, 35.40187571, 32.52161407, 45.83112306, 61.43166889, 71.11482286, 75.29755176, 74.06191255, 74.38415717,
	}

	slowk, slowd, err := ta.STOCH(high, low, closes, 5, 3, 3)
	if err != nil {
		t.Errorf(err.Error())
	}
	compare_inp_exp(t, slowk, expected_slowk, 8, "SLOWK")
	compare_inp_exp(t, slowd, expected_slowd, 8, "SLOWD")
}
