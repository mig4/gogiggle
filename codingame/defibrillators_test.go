package codingame

import (
	"encoding/csv"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestAngleDec_Radians(t *testing.T) {
	tests := []struct {
		name string
		this AngleDec
		want float64
	}{
		{"positive-1", 37.813629, 0.659972328178},
		{"positive-1", 38, 0.663225115758},
		{"negative-1", -0.127758, -0.002229797746},
		{"negative-2", -6.260310, -0.109263021696},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Radians(); !float64Eq(got, tt.want) {
				t.Errorf(
					"AngleDec(%v).Radians() = %v, want %v",
					tt.this, got, tt.want,
				)
			}
		})
	}
}

func TestNewPosition(t *testing.T) {
	type args struct {
		latitude  string
		longitude string
	}
	tests := []struct {
		name    string
		args    args
		want    Position
		wantErr bool
	}{
		{
			name:    "dot-separators",
			args:    args{"43.6071285339217", "3.87952263361082"},
			want:    Position{AngleDec(43.6071285339217), AngleDec(3.87952263361082)},
			wantErr: false,
		},
		{
			name:    "comma-separators",
			args:    args{"43,5987299452849", "3,89652239197876"},
			want:    Position{AngleDec(43.5987299452849), AngleDec(3.89652239197876)},
			wantErr: false,
		},
		{
			name:    "mixed-separators",
			args:    args{"43.6395872778854", "3,87388031141133"},
			want:    Position{AngleDec(43.6395872778854), AngleDec(3.87388031141133)},
			wantErr: false,
		},
		{
			name:    "negative-lat",
			args:    args{"-37.813629", "144.963058"},
			want:    Position{AngleDec(-37.813629), AngleDec(144.963058)},
			wantErr: false,
		},
		{
			name:    "negative-lon",
			args:    args{"51.507351", "-0.127758"},
			want:    Position{AngleDec(51.507351), AngleDec(-0.127758)},
			wantErr: false,
		},
		{
			name:    "not-a-number",
			args:    args{"nan", "nan"},
			want:    Position{},
			wantErr: true,
		},
		{
			name:    "invalid-number",
			args:    args{"wat", "invalid"},
			want:    Position{},
			wantErr: true,
		},
		{
			name:    "invalid-lat-1",
			args:    args{"-91.492013", "43.639587"},
			want:    Position{},
			wantErr: true,
		},
		{
			name:    "invalid-lat-2",
			args:    args{"100.891043", "43.639587"},
			want:    Position{},
			wantErr: true,
		},
		{
			name:    "invalid-lon-1",
			args:    args{"51.507351", "181.101287"},
			want:    Position{},
			wantErr: true,
		},
		{
			name:    "invalid-lon-2",
			args:    args{"51.507351", "-181.101287"},
			want:    Position{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPosition(tt.args.latitude, tt.args.longitude)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"NewPosition(%v) error = %v, wantErr %v",
					tt.args, err, tt.wantErr,
				)
				return
			}
			if err == nil && !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("NewPosition(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

func TestPosition_DistanceTo(t *testing.T) {
	tests := []struct {
		name  string
		base  Position
		other Position
		want  float64
	}{
		{
			name:  "case-1",
			base:  Position{43.6071285339217, 3.87952263361082},
			other: Position{43.5987299452849, 3.89652239197876},
			want:  113.94899423493638,
		},
		{
			name:  "case-2",
			base:  Position{53.349804, -6.260310},
			other: Position{53.347467, -6.309893},
			want:  315.7053162733055,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.base.DistanceTo(tt.other); !float64Eq(got, tt.want) {
				t.Errorf(
					"Position%v.DistanceTo(%v) = %v, want %v",
					tt.base, tt.other, got, tt.want,
				)
			}
		})
	}
}

func TestDefibrillatorFromFields(t *testing.T) {
	tests := []struct {
		name    string
		fields  []string
		want    Defibrillator
		wantErr bool
	}{
		{
			name: "ok-1",
			fields: []string{
				"1", "Maison de la Prevention Sante",
				"6 rue Maguelone 340000 Montpellier", "",
				"3,87952263361082", "43,6071285339217",
			},
			want: Defibrillator{
				id:       1,
				name:     "Maison de la Prevention Sante",
				address:  "6 rue Maguelone 340000 Montpellier",
				phone:    "",
				location: Position{43.6071285339217, 3.87952263361082},
			},
			wantErr: false,
		},
		{
			name: "ok-2",
			fields: []string{
				"2", "Hotel de Ville",
				"1 place Georges Freche 34267 Montpellier", "",
				"3,89652239197876", "43,5987299452849",
			},
			want: Defibrillator{
				id:       2,
				name:     "Hotel de Ville",
				address:  "1 place Georges Freche 34267 Montpellier",
				phone:    "",
				location: Position{43.5987299452849, 3.89652239197876},
			},
			wantErr: false,
		},
		{
			name: "ok-3-trim-spaces",
			fields: []string{
				"4", "  Centre municipal Garosud ",
				" 34000 Montpellier ", "04 67 34 74 62",
				"3,85859221929501", " 43,5725732056683  ",
			},
			want: Defibrillator{
				id:       4,
				name:     "Centre municipal Garosud",
				address:  "34000 Montpellier",
				phone:    "04 67 34 74 62",
				location: Position{43.5725732056683, 3.85859221929501},
			},
			wantErr: false,
		},
		{
			name: "missing-id",
			fields: []string{
				"", "Zoo de Lunaret",
				"50 avenue Agropolis 34090 Mtp", "",
				"3,87388031141133", "43,6395872778854",
			},
			want:    Defibrillator{},
			wantErr: true,
		},
		{
			name: "missing-name",
			fields: []string{
				"14", "", " 8 Avenue Louis Blanc", "04 99 58 80 31-32",
				"3,87964814275905", "43,6144971208687",
			},
			want:    Defibrillator{},
			wantErr: true,
		},
		{
			name: "invalid-position",
			fields: []string{
				"17", "Unite Service Fourriere",
				"1945 avenue de toulouse", "04 67 06 10 51",
				"130,8539", "190,5873",
			},
			want:    Defibrillator{},
			wantErr: true,
		},
		{
			name: "wrong-number-of-fields",
			fields: []string{
				"18", "cv aigoual", "Poste de police Hotel de ville",
				"789 chemin de moulares", "+33 4 34 26 69 43", "cvaigoual.com",
				"3,89399056177745", "43,5988579879724",
			},
			want:    Defibrillator{},
			wantErr: true,
		},
		{
			name:    "empty-line",
			fields:  []string{},
			want:    Defibrillator{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DefibrillatorFromFields(tt.fields)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"DefibrillatorFromLine(%v) error = %v, wantErr %v",
					tt.fields, err, tt.wantErr,
				)
				return
			}
			if err == nil && !reflect.DeepEqual(*got, tt.want) {
				t.Errorf(
					"DefibrillatorFromLine(%v) = %v, want %v",
					tt.fields, got, tt.want,
				)
			}
		})
	}
}

func TestFindDefibrillator(t *testing.T) {
	dataset := readTestDefibData(t)

	type args struct {
		defibrillators []Defibrillator
		predicate      func(*Defibrillator) bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Defibrillator
		wantErr bool
	}{
		{
			name: "find-by-name",
			args: args{dataset, func(defib *Defibrillator) bool {
				return strings.Contains(defib.name, "Batteux")
			}},
			want: &Defibrillator{
				24, "Gymnase Albert Batteux",
				"150 rue Francois-Joseph-Gossec 34000 Montpellier", "04 67 03 02 24",
				Position{43.5740760521572, 3.85685695958441},
			},
			wantErr: false,
		},
		{
			name: "find-by-address",
			args: args{dataset, func(defib *Defibrillator) bool {
				return strings.Contains(defib.address, "Metairie de Saysset")
			}},
			want: &Defibrillator{
				42, "Piscine S. BERLIOUX",
				"551 rue Metairie de Saysset MONTPELLIER", "04 67 65 38 71",
				Position{43.5904333774241, 3.89523245626307},
			},
			wantErr: false,
		},
		{
			name: "returns-first-of-multiple-matches",
			args: args{dataset, func(defib *Defibrillator) bool {
				return true
			}},
			want: &Defibrillator{
				1, "Maison de la Prevention Sante",
				"6 rue Maguelone 340000 Montpellier", "04 67 02 21 60",
				Position{43.6071285339217, 3.87952263361082},
			},
			wantErr: false,
		},
		{
			name: "no-matches",
			args: args{dataset, func(defib *Defibrillator) bool {
				return false
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindDefibrillator(tt.args.defibrillators, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindDefibrillator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDefibrillator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindNearestDefibrillator(t *testing.T) {
	dataset := readTestDefibData(t)

	wantedDefibSmallSample := Defibrillator{
		id:       1,
		name:     "Maison de la Prevention Sante",
		address:  "6 rue Maguelone 340000 Montpellier",
		phone:    "",
		location: Position{43.6071285339217, 3.87952263361082},
	}
	wantedDefibExactPosition := findDefibByName(dataset, "Cimetiere Saint-Etienne", t)
	wantedDefibFull1 := findDefibByName(dataset, "Caisse Primaire d'Assurance Maladie", t)
	wantedDefibFull2 := findDefibByName(dataset, "Amphitheatre d'O", t)

	type args struct {
		user           Position
		defibrillators []Defibrillator
	}

	tests := []struct {
		name    string
		args    args
		want    Defibrillator
		wantErr bool
	}{
		{
			name: "small-sample",
			args: args{
				user: Position{43.608177, 3.879483},
				defibrillators: []Defibrillator{
					wantedDefibSmallSample,
					Defibrillator{
						id:       2,
						name:     "Hotel de Ville",
						address:  "1 place Georges Freche 34267 Montpellier",
						phone:    "",
						location: Position{43.5987299452849, 3.89652239197876},
					},
					Defibrillator{
						id:       3,
						name:     "Zoo de Lunaret",
						address:  "50 avenue Agropolis 34090 Mtp",
						phone:    "",
						location: Position{43.6395872778854, 3.87388031141133},
					},
				},
			},
			want:    wantedDefibSmallSample,
			wantErr: false,
		},
		{
			name:    "full-dataset-exact-position",
			args:    args{Position{43.6260090150577, 3.88995587137398}, dataset},
			want:    *wantedDefibExactPosition,
			wantErr: false,
		},
		{
			name:    "full-dataset-1",
			args:    args{Position{43.606779, 3.874054}, dataset},
			want:    *wantedDefibFull1,
			wantErr: false,
		},
		{
			name:    "full-dataset-2",
			args:    args{Position{43.634646, 3.833542}, dataset},
			want:    *wantedDefibFull2,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindNearestDefibrillator(&tt.args.user, tt.args.defibrillators)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindNearestDefibrillator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("FindNearestDefibrillator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFindNearest(b *testing.B) {
	dataset := readTestDefibData(b)
	user := Position{43.608177, 3.879483}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := FindNearestDefibrillator(&user, dataset)
		if err != nil {
			b.Error(err)
		}
	}
}

/*
readTestDefibData reads in test defibrillators data file in `testdata/`
directory, parses it as CSV file, returns a slice of defibrillators created by
passing lines read from file through MapToDefibrillators.
*/
func readTestDefibData(t testing.TB) []Defibrillator {
	datafile := filepath.Join("testdata", "defibrillators-montpellier-data.csv")
	fobj, err := os.Open(datafile)
	if err != nil {
		t.Fatal(err)
	}
	defer fobj.Close()
	crd := csv.NewReader(fobj)
	crd.Comma = ';'
	crd.LazyQuotes = true
	lines, err := crd.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	defibs, err := MapToDefibrillators(lines)
	if err != nil {
		t.Fatal(err)
	}
	return defibs
}

/*
findDefibByName is a test helper that finds a Defibrillator in the given slice
with the given name (exact match) and fails the test if one is not found.
*/
func findDefibByName(
	defibrillators []Defibrillator,
	name string,
	t testing.TB,
) *Defibrillator {
	defib, err := FindDefibrillator(defibrillators, func(d *Defibrillator) bool {
		return d.name == name
	})
	if err != nil {
		t.Fatalf(
			"none of the %d defibrillators matches name %v",
			len(defibrillators), name,
		)
	}
	return defib
}

// EPSILON represents reasonable level of error for two values to still be
// considered equal.
const EPSILON = 1e-12

/*
float64Eq returns true if two floating point values are considered equal
(specifically the difference between them is smaller than EPSILON).
*/
func float64Eq(x, y float64) bool {
	return math.Abs(x-y) <= EPSILON
}
