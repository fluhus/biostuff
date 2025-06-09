package genbank

var want4 = &GenBank{
	Locus:       "HUMPROK                11613 bp    DNA     linear   PRI 07-SEP-1995",
	Definition:  "Human myotonin protein kinase (Mt-PK) gene, complete cds.",
	Accessions:  []string{"L00727"},
	Version:     "L00727.1",
	Keywords:    ".",
	Source:      "Homo sapiens (human)",
	Organism:    "Homo sapiens",
	OrganismTax: "Eukaryota; Metazoa; Chordata; Craniata; Vertebrata; Euteleostomi; Mammalia; Eutheria; Euarchontoglires; Primates; Haplorrhini; Catarrhini; Hominidae; Homo.",
	References: []map[string]string{
		{
			"":        "1  (bases 1 to 11613)",
			"AUTHORS": "Fu,Y.-H., Friedman,D.L., Richards,S., Pearlman,J.A., Gibbs,R.A., Pizzuti,A., Ashizawa,T., Perryman,M.B., Scarlato,G., Fenwick,R.G. Jr. and Caskey,C.T.",
			"TITLE":   "Decreased expression of myotonin-protein kinase messenger RNA and protein in adult form of myotonic dystrophy",
			"JOURNAL": "Science 260 (5105), 235-238 (1993)",
			"PUBMED":  "8469976",
		},
		{
			"":        "2  (sites)",
			"AUTHORS": "Fu,Y.-H., Pizzuti,A., Fenwick,R.G. Jr., King,J., Rajnarayan,S., Dunne,P.W., Dubel,J., Nasser,G.A., Ashizawa,T., de Jong,P., Wieringa,B., Korneluk,R., Perryman,M.B., Epstein,H.F. and Caskey,C.T.",
			"TITLE":   "An unstable triplet repeat in a gene related to myotonic muscular dystrophy",
			"JOURNAL": "Science 255 (5049), 1256-1258 (1992)",
			"PUBMED":  "1546326",
		},
	},
	Features: []*Feature{
		{
			Type: "source",
			Fields: map[string]string{
				"":           "1..11613",
				"chromosome": "19",
				"clone":      "cosmid MDY1",
				"db_xref":    "taxon:9606",
				"map":        "19q13.2",
				"mol_type":   "genomic DNA",
				"organism":   "Homo sapiens",
				"tissue_lib": "St. Louis",
			},
		},
		{
			Type: "gene",
			Fields: map[string]string{
				"":     "1310..11188",
				"gene": "Mt-PK",
			},
		},
		{
			Type: "mRNA",
			Fields: map[string]string{
				"":     "join(1310..2135,2391..2474,2554..2649,3272..3420, 3693..3786,4040..4246,4322..4585,5901..5986,8184..8295, 9008..9165,9271..9368,9540..9586,9876..9965,10296..11188)",
				"gene": "Mt-PK",
				"note": "alternatively spliced Form I mRNA; exons were positioned according to M87312; I and III are the major forms in most tissues except fetal and newborn muscle",
			},
		},
		{
			Type: "exon",
			Fields: map[string]string{
				"":     "1310..2135",
				"gene": "Mt-PK",
				"note": "Form I,V,VI,VII,VIII mRNA",
			},
		},
		{
			Type: "CDS",
			Fields: map[string]string{
				"":            "join(1854..2135,2391..2474,2554..2649,3272..3420, 3693..3786,4040..4246,4322..4585,5901..5986,8184..8295, 9008..9165,9271..9368,9540..9586,9876..9965,10296..10448)",
				"codon_start": "1",
				"gene":        "Mt-PK",
				"note":        "Form I mRNA",
				"product":     "myotonin-protein kinase, Form I",
				"protein_id":  "AAA75236.1",
				"translation": "MGGHFWPPEPYTVFMWGSPWEADSPRVKLRGREKGRQTEGGAFPLVSSALSGDPRFFSPTTPPAEPIVVRLKEVRLQRDDFEILKVIGRGAFSEVAVVKMKQTGQVYAMKIMNKWDMLKRGEVSCFREERDVLVNGDRRWITQLHFAFQDENYLYLVMEYYVGGDLLTLLSKFGERIPAEMARFYLAEIVMAIDSVHRLGYVHRDIKPDNILLDRCGHIRLADFGSCLKLRADGTVRSLVAVGTPDYLSPEILQAVGGGPGTGSYGPECDWWALGVFAYEMFYGQTPFYADSTAETYGKIVHYKEHLSLPLVDEGVPEEARDFIQRLLCPPETRLGRGGAGDFRTHPFFFGLDWDGLRDSVPPFTPDFEGATDTCNFDLVEDGLTAMVSGGGETLSDIREGAPLGVHLPFVGYSYSCMALRDSEVPGPTPMEVEAEQLLEPHVQAPSLEPSVSPQDETAEVAVPAAVPAAEAEAEVTLRELQEALEEEVLTRQSLSREMEAIRTDNQNFASQLREAEARNRDLEAHVRQLQERMELLQAEGATAVTGVPSPRATDPPSHLDGPPAVAVGQCPLVGPGPMHRRHLLLPARVPRPGLSEALSLLLFAVVLSRAAALGCIGLVAHAGQLTAVWRRPGAARAP",
			},
		},
		{
			Type: "exon",
			Fields: map[string]string{
				"":     "2044..2135",
				"gene": "Mt-PK",
				"note": "Form II,III,and IV mRNA",
			},
		},
		{
			Type: "CDS",
			Fields: map[string]string{
				"":            "join(2406..2474,2554..2649,3272..3420,3693..3786, 4040..4246,4322..4585,5901..5986,8184..8295,9008..9165, 9271..9368,9540..9586,9876..9965,10296..10448)",
				"codon_start": "1",
				"gene":        "Mt-PK",
				"note":        "Form II,III,IV mRNA products each are missing the 99 residues from the N-terminus of Form I",
				"product":     "myotonin-protein kinase, Form II,III,IV",
				"protein_id":  "AAA75240.1",
				"translation": "MKQTGQVYAMKIMNKWDMLKRGEVSCFREERDVLVNGDRRWITQLHFAFQDENYLYLVMEYYVGGDLLTLLSKFGERIPAEMARFYLAEIVMAIDSVHRLGYVHRDIKPDNILLDRCGHIRLADFGSCLKLRADGTVRSLVAVGTPDYLSPEILQAVGGGPGTGSYGPECDWWALGVFAYEMFYGQTPFYADSTAETYGKIVHYKEHLSLPLVDEGVPEEARDFIQRLLCPPETRLGRGGAGDFRTHPFFFGLDWDGLRDSVPPFTPDFEGATDTCNFDLVEDGLTAMVSGGGETLSDIREGAPLGVHLPFVGYSYSCMALRDSEVPGPTPMEVEAEQLLEPHVQAPSLEPSVSPQDETAEVAVPAAVPAAEAEAEVTLRELQEALEEEVLTRQSLSREMEAIRTDNQNFASQLREAEARNRDLEAHVRQLQERMELLQAEGATAVTGVPSPRATDPPSHLDGPPAVAVGQCPLVGPGPMHRRHLLLPARVPRPGLSEALSLLLFAVVLSRAAALGCIGLVAHAGQLTAVWRRPGAARAP",
			},
		},
		{
			Type: "exon",
			Fields: map[string]string{
				"":     "2554..2649",
				"gene": "Mt-PK",
				"note": "Form I,II,III,IV,V,VI,VII,VIII mRNA",
			},
		},
		{
			Type: "repeat_region",
			Fields: map[string]string{
				"":           "5153..5434",
				"gene":       "Mt-PK",
				"rpt_family": "Alu-J",
				"rpt_type":   "dispersed",
			},
		},
		{
			Type: "exon",
			Fields: map[string]string{
				"":     "5901..5986",
				"gene": "Mt-PK",
				"note": "Form I,II,III,IV,V,VI,VII,VIII mRNA",
			},
		},
		{
			Type: "repeat_region",
			Fields: map[string]string{
				"":         "6363..6373",
				"gene":     "Mt-PK",
				"rpt_type": "direct",
			},
		},
		{
			Type: "misc_feature",
			Fields: map[string]string{
				"":     "7715..8045",
				"gene": "Mt-PK",
				"note": "similar to EST sequence in GenBank Accession Number H04589",
			},
		},
		{
			Type: "exon",
			Fields: map[string]string{
				"":     "8184..8295",
				"gene": "Mt-PK",
				"note": "Form I,II,III,IV,V,VI,VII,VIII mRNA",
			},
		},
		{
			Type: "misc_feature",
			Fields: map[string]string{
				"":     "9571..10953",
				"gene": "Mt-PK",
				"note": "region previously sequenced; see GenBank Accession Number M87313",
			},
		},
		{
			Type: "exon",
			Fields: map[string]string{
				"":     "9876..9965",
				"gene": "Mt-PK",
				"note": "Form I,II,III,IV,V,VI mRNA",
			},
		},
		{
			Type: "repeat_region",
			Fields: map[string]string{
				"":             "10677..10703",
				"gene":         "Mt-PK",
				"note":         "9 gct repeats in this region; expanded variants known to be present in patients with myotonic dystrophy",
				"rpt_unit_seq": "gct",
			},
		},
	},
	Origin: origin4,
}

// Taken from:
// https://www.ncbi.nlm.nih.gov/nucleotide/L00727
//
// Some features were removed to make the manual testing easier.
const input4 = `LOCUS       HUMPROK                11613 bp    DNA     linear   PRI 07-SEP-1995
DEFINITION  Human myotonin protein kinase (Mt-PK) gene, complete cds.
ACCESSION   L00727
VERSION     L00727.1
KEYWORDS    .
SOURCE      Homo sapiens (human)
  ORGANISM  Homo sapiens
            Eukaryota; Metazoa; Chordata; Craniata; Vertebrata; Euteleostomi;
            Mammalia; Eutheria; Euarchontoglires; Primates; Haplorrhini;
            Catarrhini; Hominidae; Homo.
REFERENCE   1  (bases 1 to 11613)
  AUTHORS   Fu,Y.-H., Friedman,D.L., Richards,S., Pearlman,J.A., Gibbs,R.A.,
            Pizzuti,A., Ashizawa,T., Perryman,M.B., Scarlato,G., Fenwick,R.G.
            Jr. and Caskey,C.T.
  TITLE     Decreased expression of myotonin-protein kinase messenger RNA and
            protein in adult form of myotonic dystrophy
  JOURNAL   Science 260 (5105), 235-238 (1993)
   PUBMED   8469976
REFERENCE   2  (sites)
  AUTHORS   Fu,Y.-H., Pizzuti,A., Fenwick,R.G. Jr., King,J., Rajnarayan,S.,
            Dunne,P.W., Dubel,J., Nasser,G.A., Ashizawa,T., de Jong,P.,
            Wieringa,B., Korneluk,R., Perryman,M.B., Epstein,H.F. and
            Caskey,C.T.
  TITLE     An unstable triplet repeat in a gene related to myotonic muscular
            dystrophy
  JOURNAL   Science 255 (5049), 1256-1258 (1992)
   PUBMED   1546326
FEATURES             Location/Qualifiers
     source          1..11613
                     /organism="Homo sapiens"
                     /mol_type="genomic DNA"
                     /db_xref="taxon:9606"
                     /chromosome="19"
                     /map="19q13.2"
                     /clone="cosmid MDY1"
                     /tissue_lib="St. Louis"
     gene            1310..11188
                     /gene="Mt-PK"
     mRNA            join(1310..2135,2391..2474,2554..2649,3272..3420,
                     3693..3786,4040..4246,4322..4585,5901..5986,8184..8295,
                     9008..9165,9271..9368,9540..9586,9876..9965,10296..11188)
                     /gene="Mt-PK"
                     /note="alternatively spliced Form I mRNA; exons were
                     positioned according to M87312; I and III are the major
                     forms in most tissues except fetal and newborn muscle"
     exon            1310..2135
                     /gene="Mt-PK"
                     /note="Form I,V,VI,VII,VIII mRNA"
     CDS             join(1854..2135,2391..2474,2554..2649,3272..3420,
                     3693..3786,4040..4246,4322..4585,5901..5986,8184..8295,
                     9008..9165,9271..9368,9540..9586,9876..9965,10296..10448)
                     /gene="Mt-PK"
                     /note="Form I mRNA"
                     /codon_start=1
                     /product="myotonin-protein kinase, Form I"
                     /protein_id="AAA75236.1"
                     /translation="MGGHFWPPEPYTVFMWGSPWEADSPRVKLRGREKGRQTEGGAFP
                     LVSSALSGDPRFFSPTTPPAEPIVVRLKEVRLQRDDFEILKVIGRGAFSEVAVVKMKQ
                     TGQVYAMKIMNKWDMLKRGEVSCFREERDVLVNGDRRWITQLHFAFQDENYLYLVMEY
                     YVGGDLLTLLSKFGERIPAEMARFYLAEIVMAIDSVHRLGYVHRDIKPDNILLDRCGH
                     IRLADFGSCLKLRADGTVRSLVAVGTPDYLSPEILQAVGGGPGTGSYGPECDWWALGV
                     FAYEMFYGQTPFYADSTAETYGKIVHYKEHLSLPLVDEGVPEEARDFIQRLLCPPETR
                     LGRGGAGDFRTHPFFFGLDWDGLRDSVPPFTPDFEGATDTCNFDLVEDGLTAMVSGGG
                     ETLSDIREGAPLGVHLPFVGYSYSCMALRDSEVPGPTPMEVEAEQLLEPHVQAPSLEP
                     SVSPQDETAEVAVPAAVPAAEAEAEVTLRELQEALEEEVLTRQSLSREMEAIRTDNQN
                     FASQLREAEARNRDLEAHVRQLQERMELLQAEGATAVTGVPSPRATDPPSHLDGPPAV
                     AVGQCPLVGPGPMHRRHLLLPARVPRPGLSEALSLLLFAVVLSRAAALGCIGLVAHAG
                     QLTAVWRRPGAARAP"
     exon            2044..2135
                     /gene="Mt-PK"
                     /note="Form II,III,and IV mRNA"
     CDS             join(2406..2474,2554..2649,3272..3420,3693..3786,
                     4040..4246,4322..4585,5901..5986,8184..8295,9008..9165,
                     9271..9368,9540..9586,9876..9965,10296..10448)
                     /gene="Mt-PK"
                     /note="Form II,III,IV mRNA products each are missing the
                     99 residues from the N-terminus of Form I"
                     /codon_start=1
                     /product="myotonin-protein kinase, Form II,III,IV"
                     /protein_id="AAA75240.1"
                     /translation="MKQTGQVYAMKIMNKWDMLKRGEVSCFREERDVLVNGDRRWITQ
                     LHFAFQDENYLYLVMEYYVGGDLLTLLSKFGERIPAEMARFYLAEIVMAIDSVHRLGY
                     VHRDIKPDNILLDRCGHIRLADFGSCLKLRADGTVRSLVAVGTPDYLSPEILQAVGGG
                     PGTGSYGPECDWWALGVFAYEMFYGQTPFYADSTAETYGKIVHYKEHLSLPLVDEGVP
                     EEARDFIQRLLCPPETRLGRGGAGDFRTHPFFFGLDWDGLRDSVPPFTPDFEGATDTC
                     NFDLVEDGLTAMVSGGGETLSDIREGAPLGVHLPFVGYSYSCMALRDSEVPGPTPMEV
                     EAEQLLEPHVQAPSLEPSVSPQDETAEVAVPAAVPAAEAEAEVTLRELQEALEEEVLT
                     RQSLSREMEAIRTDNQNFASQLREAEARNRDLEAHVRQLQERMELLQAEGATAVTGVP
                     SPRATDPPSHLDGPPAVAVGQCPLVGPGPMHRRHLLLPARVPRPGLSEALSLLLFAVV
                     LSRAAALGCIGLVAHAGQLTAVWRRPGAARAP"
     exon            2554..2649
                     /gene="Mt-PK"
                     /note="Form I,II,III,IV,V,VI,VII,VIII mRNA"
     repeat_region   5153..5434
                     /gene="Mt-PK"
                     /rpt_family="Alu-J"
                     /rpt_type=dispersed
     exon            5901..5986
                     /gene="Mt-PK"
                     /note="Form I,II,III,IV,V,VI,VII,VIII mRNA"
     repeat_region   6363..6373
                     /gene="Mt-PK"
                     /rpt_type=direct
     misc_feature    7715..8045
                     /gene="Mt-PK"
                     /note="similar to EST sequence in GenBank Accession Number
                     H04589"
     exon            8184..8295
                     /gene="Mt-PK"
                     /note="Form I,II,III,IV,V,VI,VII,VIII mRNA"
     misc_feature    9571..10953
                     /gene="Mt-PK"
                     /note="region previously sequenced; see GenBank Accession
                     Number M87313"
     exon            9876..9965
                     /gene="Mt-PK"
                     /note="Form I,II,III,IV,V,VI mRNA"
     repeat_region   10677..10703
                     /gene="Mt-PK"
                     /note="9 gct repeats in this region; expanded variants
                     known to be present in patients with myotonic dystrophy"
                     /rpt_unit_seq="gct"
ORIGIN      
        1 ccatggcctc tctgcacccc gcctcagggt cagggtcagg gtcatgctgg gagctccctc
       61 tcctaggacc ctccccccaa aagtgggctc tatggccctc tcccctggtt tcctgtggcc
      121 tggggcaagc caggagggcc agcatggggc agctgccagg ggcgcagccg acaggcaggt
      181 gttcggcgcc agcctctcca gctgccccaa caggtgccca ggcgctggga gggcggtgac
      241 tcacgcgggc cctgtgggag aaccagcttt gcagacaggc gccaccagtg ccccctcctc
      301 tgcgatccag gagggacaac tttgggttct tctgggtgtg tctccttctt ttgtaggttc
      361 tgcacccacc cccaccccca gccccaaagt ctcggttcct atgagccgtg tgggtcagcc
      421 accattcccg ccaccccggg tccctgcgtc ctttagttct cctggcccag ggcctccaac
      481 cttccagctg tcccacaaaa ccccttcttg caagggcttt ccagggcctg gggccagggc
      541 tggaaggagg atgcttccgc ttctgccagc tgccttgtct gcccacctcc tccccaagcc
      601 caggactcgg gctcactggt cactggtttc tttcattccc agcaccctgc tcctctggcc
      661 ctcatatgtc tggccctcag tgactggtgt ttggtttttg gcctgtgtgt aacaaactgt
      721 gtgtgacact tgtttcctgt ttctccgcct tcccctgctt cctcttgtgt ccatctcttt
      781 ctgacccagg cctggttcct ttccctcctc ctcccatttc acagatggga aggtggcggc
      841 caagaagggc caggccattc agcctctgga aaaaccttct cccaacctcc cacagcccct
      901 aatgactctc ctggcctccc tttagtagag gatgaagttg ggttggcagg gtaaactgag
      961 accgggtggg gtaggggtct ggcgctcccg ggaggagcac tccttttgtg gcccgagctg
     1021 catctcgcgg cccctcccct gcaaggcctg gggcggggga gggggccagg gttcctgctg
     1081 ccttaaaagg gctcaatgtc ttggctctct cctccctccc ccgtcctcag ccctggctgg
     1141 ttcgtccctg ctggcccact ctcccggaac cccccggaac ccctctcttt cctccagaac
     1201 ccactgtctc ctctccttcc ctcccctccc atacccatcc ctctctccat cctgcctcca
     1261 cttcttccac ccccgggagt ccaggcctcc ctgtccccac agtccctgag ccacaagcct
     1321 ccaccccagc tggtccccca cccaggctgc ccagtttaac attcctagtc ataggacctt
     1381 gacttctgag aggcctgatt gtcatctgta aataaggggt aggactaaag cactcctcct
     1441 ggaggactga gagatgggct ggaccggagc acttgagtct gggatatgtg accatgctac
     1501 ctttgtctcc ctgtcctgtt ccttccccca gccccaaatc cagggttttc caaagtgtgg
     1561 ttcaagaacc acctgcatct gaatctagag gtactggata caaccccacg tctgggccgt
     1621 tacccaggac attctacatg agaacgtggg ggtggggccc tggctgcacc tgaactgtca
     1681 cctggagtca gggtggaagg tggaagaact gggtcttatt tccttctccc cttgttcttt
     1741 agggtctgtc cttctgcaga ctccgttacc ccaccctaac catcctgcac acccttggag
     1801 ccctctgggc caatgccctg tcccgcaaag ggcttctcag gcatctcacc tctatgggag
     1861 ggcatttttg gcccccagaa ccttacacgg tgtttatgtg gggaagcccc tgggaagcag
     1921 acagtcctag ggtgaagctg agaggcagag agaaggggag acagacagag ggtggggctt
     1981 tcccccttgt ctccagtgcc ctttctggtg accctcggtt cttttccccc accacccccc
     2041 cagcggagcc catcgtggtg aggcttaagg aggtccgact gcagagggac gacttcgaga
     2101 ttctgaaggt gatcggacgc ggggcgttca gcgaggtaag ccgaaccggg cgggagcctg
     2161 acttgactcg tggtgggcgg ggcatagggg ttggggcggg gccttagaaa ttgatgaatg
     2221 accgagcctt agaacctagg gctgggctgg aggcggggct tgggaccaat gggcgtggtg
     2281 tggcaggtgg ggcggggcca cggctgggtg cagaagcggg tggagttggg tctgggcgag
     2341 cccttttgtt ttcccgccgt ctccactctg tctcactatc tcgacctcag gtagcggtag
     2401 tgaagatgaa gcagacgggc caggtgtatg ccatgaagat catgaacaag tgggacatgc
     2461 tgaagagggg cgaggtgagg ggctgggcgg acgtgggggg ctttgaggat ccgcgccccg
     2521 tctccggctg cagctcctcc gggtgccctg caggtgtcgt gcttccgtga ggagagggac
     2581 gtgttggtga atggggaccg gcggtggatc acgcagctgc acttcgcctt ccaggatgag
     2641 aactacctgg tgagctccgg gccgggggga ctaggaagag ggacaagagc ccgtgctgtc
     2701 actggacgag gaggtgggga gaggaagctc taggattggg ggtgctgccc ggaaacgtct
     2761 gtgggaaagt ctgtgtgcgg taagagggtg tgtcaggtgg atgaggggcc ttccctatct
     2821 gagacgggga tggtgtcctt cactgcccgt ttctggggtg atctggggga ctcttataaa
     2881 gatgtctctg ttgcgggggg tctcttacct ggaatgggat aggtcttcag gaattctaac
     2941 ggggccactg cctagggaag gagtgtctgg gacctattct ctgggtgttg ggtggcctct
     3001 gggttctctt tcccagaaca tctcaggggg agtgaatctg cccagtgaca tcccaggaaa
     3061 gtttttttgt ttgtgttttt ttttgagggg cgggggcggg ggccgcaggt ggtctctgat
     3121 ttggcccggc agatctctat ggttatctct gggctggggc tgcaggtctc tgcccaagga
     3181 tggggtgtct ctgggagggg ttgtcccagc catccgtgat ggatcagggc ctcaggggac
     3241 taccaaccac ccatgacgaa ccccttctca gtacctggtc atggagtatt acgtgggcgg
     3301 ggacctgctg acactgctga gcaagtttgg ggagcggatt ccggccgaga tggcgcgctt
     3361 ctacctggcg gagattgtca tggccataga ctcggtgcac cggcttggct acgtgcacag
     3421 gtgggcgcag catggccgag gggatagcaa gcttgttccc tggccgggtt cttggaaggt
     3481 cagagcccag agaggccagg gcctggagag ggaccttctt ggttggggcc caccgggggg
     3541 tgcctgggag taggggtcag aactgtagaa gccctacagg ggcggaaccc gaggaagtgg
     3601 ggtcccaggt ggcactgccc ggaggggcgg agcctggtgg gaccacagaa gggaggttca
     3661 tttatcccac ccttctcttt tcctccgtgc agggacatca aacccgacaa catcctgctg
     3721 gaccgctgtg gccacatccg cctggccgac ttcggctctt gcctcaagct gcgggcagat
     3781 ggaacggtga gccagtgccc tggccacaga gcaactgggg ctgctgatga gggatggaag
     3841 gcacagagtg tgggagcggg actggatttg gaggggaaaa gaggtggtgt gacccaggct
     3901 taagtgtgca tctgtgtggc ggagtattag accaggcaga gggaggggct aagcatttgg
     3961 ggagtggttg gaaggagggc ccagagctgg tgggcccaga ggggtgggcc caagcctcgc
     4021 tctgctcctt ttggtccagg tgcggtcgct ggtggctgtg ggcaccccag actacctgtc
     4081 ccccgagatc ctgcaggctg tgggcggtgg gcctgggaca ggcagctacg ggcccgagtg
     4141 tgactggtgg gcgctgggtg tattcgccta tgaaatgttc tatgggcaga cgcccttcta
     4201 cgcggattcc acggcggaga cctatggcaa gatcgtccac tacaaggtga gcacggccgc
     4261 agggagacct ggcctctccc ggtaggcgct cccagctatc gcctcctctc cctctgagca
     4321 ggagcacctc tctctgccgc tggtggacga aggggtccct gaggaggctc gagacttcat
     4381 tcagcggttg ctgtgtcccc cggagacacg gctgggccgg ggtggagcag gcgacttccg
     4441 gacacatccc ttcttctttg gcctcgactg ggatggtctc cgggacagcg tgcccccctt
     4501 tacaccggat ttcgaaggtg ccaccgacac atgcaacttc gacttggtgg aggacgggct
     4561 cactgccatg gtgagcgggg gcggggtagg tacctgtggc ccctgctcgg ctgcgggaac
     4621 ctccccatgc tccctccata aagttggagt aaggacagtg cctaccttct ggggtcctga
     4681 atcactcatt ccccagagca cctgctctgt gcccatctac tactgaggac ccagcagtga
     4741 cctagactta cagtccagtg ggggaacaca gagcagtctt cagacagtaa ggccccagag
     4801 tgatcagggc tgagacaatg gagtgcaggg ggtgggggac tcctgactca gcaaggaagg
     4861 tcctggaggg ctttctggag tggggagcta tctgagctga gacttggagg gatgagaagc
     4921 aggagaggac tcctcctccc ttaggccgtc tctcttcacc gtgtaacaag ctgtcatggc
     4981 atgcttgctc ggctctgggt gcccttttgc tgaacaatac tggggatcca gcacggacca
     5041 gatgagctct ggtccctgcc ctcatccagt tgcagtctag agaattagag aattatggag
     5101 agtgtggcag gtgccctgaa gggaagcaac aggatacaag aaaaaatgat ggggccaggc
     5161 acggtgctca cgcctgtaac cccagcaatt tggcaggccg aagtgggtgg attgcttgag
     5221 cccaggagtt cgagaccagc ctgggcaatg tggtgagacc cccgtctcta caaaaatgtt
     5281 ttaaaaattg gttgggcgtg gtggcgcatg cctgtatact cagctactag ggtggccgac
     5341 gtgggcttga gcccaggagg tcaaggctgc agtgagctgt gattgtgcca ctgcactcca
     5401 gcctgggcaa cggagagaga ctctgtctca aaaataagat aaactgaaat taaaaaatag
     5461 gctgggctgg ccgggcgtgg tggctcacgc ctgtaatctc agcactttgg gaggccgagg
     5521 cgggtggatc acgaggtcag aagatggaga ccagcctggc cagcgtggcg aaaccccgtc
     5581 tctaccaaaa atataaaaaa ttagccaggc gtggtagagg gcgcctgtaa tctcagctac
     5641 tcaggacgct gaggcaggag aatcgcctga acctgggagg cggaggttgc agtgagctga
     5701 gattgcacca ctgcactcca gcctgggtaa cagagcgaga ctccgtatca aagaaaaaga
     5761 aaaaagaaaa aatgctggag gggccacttt agataagccc tgagttgggg ctggtttggg
     5821 gggaacatgt aagccaagat caaaaagcag tgaggggccc gccctgacga ctgctgctca
     5881 catctgtgtg tcttgcgcag gagacactgt cggacattcg ggaaggtgcg ccgctagggg
     5941 tccacctgcc ttttgtgggc tactcctact cctgcatggc cctcaggtaa gcactgccct
     6001 ggacggcctc caggggccac gaggctgctt gagcttcctg ggtcctgctc cttggcagcc
     6061 aatggagttg caggatcagt cttggaacct tactgttttg ggcccaaaga ctcctaagag
     6121 gccagagttg gaggacctta aattttcaga tctatgtact tcaaaatgtt agattgaatt
     6181 ttaaaacctc agagtcacag actgggcttc ccagaatctt gtaaccatta acttttacgt
     6241 ctgtagtaca cagagccaca ggacttcaga acttggaaaa tatgaagttt agacttttac
     6301 aatcagttgt aaaagaatgc aaattctttg aatcagccat ataacaataa ggccatttaa
     6361 aagtattaat ttaggcgggc cgcggtggct cacgcctgta atcctagcac tttgggaggc
     6421 caaggcaggt ggatcatgag gtcaggagat cgagaccatc ctggctaaca cggtgaaacc
     6481 ccgtctctac taaaaataca aaaaaattag ccgggcatgg tggcgggcgc ttgcggtccc
     6541 agctacttgg gaggcgaggc aggagaatgg catgaacccg ggaggcggag cttgcagtga
     6601 gccgagatca tgccactgca ctccagcctg ggcgacagag caagactccg tctcaaaaaa
     6661 aaaaaaaaaa aaagtattta tttaggccgg gtgtggtggc tcacgcctgt aattccagtg
     6721 ctttgggagg atgaggtggg tggatcacct gaggtcagga gttcgagacc agcctgacca
     6781 acgtggagaa acctcatctc tactaaaaaa caaaattagc caggcatggt ggcatatacc
     6841 tgtaatccca gctactcagg aggctgaggc aggagaatca gaacccagga gggggaggtt
     6901 gtggttagct gagatcgtgc cattgcattc cagcctgggc aacaagagtg aaacttcatc
     6961 tcaaaaaaaa aaaaaaaaaa gtactaattt acaggctggg catggtggct cacgcttgga
     7021 atcccagcac tttgggaggc tgaagtggac ggattgcttc agcccaggag ttcaagacca
     7081 gcctgagcaa cataatgaga ccctgtctct acaaaaaatt gaaaaaatcg tgccaggcat
     7141 ggtggtctgt gcctgcagtc ctagctactc aggagtctga agtaggagaa tcacttgagc
     7201 ctggagtttg aggcttcagt gagccatgat agattccagc ctaggcaaca aagtgagacc
     7261 tggtctcaac aaaagtatta attacacaaa taatgcattg cttatcacaa gtaaattaga
     7321 aaatacagat aaggaaaagg aagttgatat ctcgtgagct caccagatgg cagtggtccc
     7381 tggctcacac gtgtactgac acatgtttaa atagtggaga acaggtgttt ttttggtttg
     7441 tttttttccc cttcctcatg ctactttgtc taagagaaca gttggttttc tagtcagctt
     7501 ttattactgg acaacattac acatactata ccttatcatt aatgaactcc agcttgattc
     7561 tgaaccgctg cggggcctga acggtgggtc aggattgaac ccatcctcta ttagaaccca
     7621 ggcgcatgtc caggatagct aggtcctgag ccgtgttccc acaggaggga ctgctgggtt
     7681 ggaggggaca gccacttcat accccaggga ggagctgtcc ccttcccaca gctgagtggg
     7741 gtgtgctgac ctcaagttgc catcttgggg tcccatgccc agtcttagga ccacatctgt
     7801 ggaggtggcc agagccaagc agtctcccca tcaggtcggc ctccctgtcc tgaggccctg
     7861 agaagagggg tctgcagcgg tcacatgtca agggaggaga tgagctgacc ctagaacatg
     7921 ggggtctgga ccccaagtcc ctgcagaagg tttagaaaga gcagctccca ggggcccaag
     7981 gccaggagag gggcagggct tttcctaagc agaggagggg ctattggcct acctgggact
     8041 ctgttctctt cgctctgctg ctccccttcc tcaaatcagg aggtcttgga agcagctgcc
     8101 cctacccaca ggccagaagt tctggttctc caccagataa tcagcattct gtctccctcc
     8161 ccactccctc ctcctctccc cagggacagt gaggtcccag gccccacacc catggaagtg
     8221 gaggccgagc agctgcttga gccacacgtg caagcgccca gcctggagcc ctcggtgtcc
     8281 ccacaggatg aaacagtaag ttggtggagg ggagggggtc cgtcagggac aattgggaga
     8341 gaaaaggtga gggcttcccg ggtggcgtgc actgtagagc cctctaggga cttcctgaac
     8401 agaagcagac agaaaccacg gagagacgag gttacttcag acatgggacg gtctctgtag
     8461 ttacagtggg gcattaagta agggtgtgtg tgttgctggg gatctgagaa gtcgatcttt
     8521 gagctgagcg ctggtgaagg agaaacaagc catggaagga aaggtgccaa gtggtcaggc
     8581 gagagcctcc agggcaaagg ccttgggcag gtgggaatcc tgatttgttc ctgaaaggta
     8641 gtttggctga atcattcctg agaaggctgg agaggccagc aggaaacaaa acccagcaag
     8701 gccttttgtc gtgagggcat tagggagctg gagggatttt gagcagcaga gggacatagg
     8761 ttgtgttagt gtttgagcac cagccctctg gtccctgtgt agatttagag gaccagactc
     8821 agggatgggg ctgagggagg tagggaaggg agggggcttg gatcattgca ggagctatgg
     8881 ggattccaga aatgttgagg ggacggagga gtaggggata aacaaggatt cctagcctgg
     8941 aaccagtgcc caagtcctga gtcttccagg agccacaggc agccttaagc ctggtcccca
     9001 tacacaggct gaagtggcag ttccagcggc tgtccctgcg gcagaggctg aggccgaggt
     9061 gacgctgcgg gagctccagg aagccctgga ggaggaggtg ctcacccggc agagcctgag
     9121 ccgggagatg gaggccatcc gcacggacaa ccagaacttc gccaggtcgg gatcggggcc
     9181 ggggccgggg ccgggatgcg ggccggtggc aacccttggc agcccctctc gtccggcccg
     9241 gacggactca ccgtccttac ctccccacag tcaactacgc gaggcagagg ctcggaaccg
     9301 ggacctagag gcacacgtcc ggcagttgca ggagcggatg gagttgctgc aggcagaggg
     9361 agccacaggt gagtccctca tgtgtcccct tccccggagg accgggagga ggtgggccgt
     9421 ctgctccgcg gggcgtgtat agacacctgg aggagggaag ggacccacgc tggggcacgc
     9481 cgcgccaccg ccctccttcg cccctccacg cgccctatgc ctctttcttc tccttccagc
     9541 tgtcacgggg gtccccagtc cccgggccac ggatccacct tcccatgtaa gacccctctc
     9601 tttcccctgc ctcagacctg ctgcccattc tgcagatccc ctccctggct cctggtctcc
     9661 ccgtccagat atagggctca ccctacgtct ttgcgacttt agagggcaga agccctttat
     9721 tcagccccag atctccctcc gttcaggcct caccagattc cctccgggat ctccctagat
     9781 aacctcccca acctcgattc cgctcgctgt ctctcgcccc accgctgagg gctgggctgg
     9841 gctccgatcg ggtcacctgt cccttctctc tccagctaga tggccccccg gccgtggctg
     9901 tgggccagtg cccgctggtg gggccaggcc ccatgcaccg ccgccacctg ctgctccctg
     9961 ccagggtacg tccggctgcc cacgcccccc tccgccgtcg cgccccgcgc tccacccgcc
    10021 ccgtgccacc cgcttagctg cgcatttgcg gggctgggcc cacggcagga gggcggatct
    10081 tcgggcagcc aatcaacaca ggccgctagg aagcagccaa tgacgagttc ggacgggatt
    10141 cgaggcgtgc gagtggacta acaacagctg taggctgttg gggcgggggc ggggcgcagg
    10201 gaagagtgcg ggcccaccta tgggcgtagg cggggcgagt cccaggagcc aatcagaggc
    10261 ccatgccggg tgttgacctc gccctctccc cgcaggtccc taggcctggc ctatcggagg
    10321 cgctttccct gctcctgttc gccgttgttc tgtctcgtgc cgccgccctg ggctgcattg
    10381 ggttggtggc ccacgccggc caactcaccg cagtctggcg ccgcccagga gccgcccgcg
    10441 ctccctgaac cctagaactg tcttcgactc cggggccccg ttggaagact gagtgcccgg
    10501 ggcacggcac agaagccgcg cccaccgcct gccagttcac aaccgctccg agcgtgggtc
    10561 tccgcccagc tccagtcctg tgatccgggc ccgcccccta gcggccgggg agggaggggc
    10621 cgggtccgcg gccggcgaac ggggctcgaa gggtccttgt agccgggaat gctgctgctg
    10681 ctgctgctgc tgctgctgct gctgggggga tcacagacca tttctttctt tcggccaggc
    10741 tgaggccctg acgtggatgg gcaaactgca ggcctgggaa ggcagcaagc cgggccgtcc
    10801 gtgttccatc ctccacgcac ccccacctat cgttggttcg caaagtgcaa agctttcttg
    10861 tgcatgacgc cctgctctgg ggagcgtctg gcgcgatctc tgcctgctta ctcgggaaat
    10921 ttgcttttgc caaacccgct ttttcgggga tcccgcgccc ccctcctcac ttgcgctgct
    10981 ctcggagccc cagccggctc cgcccgcttc ggcggtttgg atatttattg acctcgtcct
    11041 ccgactcgct gacaggctac aggaccccca acaaccccaa tccacgtttt ggatgcactg
    11101 agaccccgac attcctcggt atttattgtc tgtccccacc taggaccccc acccccgacc
    11161 ctcgcgaata aaaggccctc catctgccca aagctctgga ctccacagtg tccgcggttt
    11221 gcgttgtggg ccggagctcc gcagcgggcc aatccggagg cgtgtggagg cggccgaagg
    11281 tctgggagga gctagcggga tgcgaagcgg ccgaatcagg gttgggggag gaaaagccac
    11341 ggggcggggc tttggcgtcc ggccaatagg agggcgagcg ggccacccgg aggcaccgcc
    11401 cccgcccagc tgtggcccag ctgtgccacc gagcgtcgag aagagggggc tgggctggca
    11461 gcgcgcgcgg ccatcctcct tccactgcgc ctgcgcacgc cacgcgcatc cgctcctggg
    11521 acgcaagctc gagaaaagtt gctgcaaact ttctagcccg ttccccgccc ctcctcccgg
    11581 ccagacccgc cccccctgcg gagccgggaa ttc
//`

const origin4 = "ccatggcctctctgcaccccgcctcagggtcagggtcagggtcatgctgggagctccctc" +
	"tcctaggaccctccccccaaaagtgggctctatggccctctcccctggtttcctgtggcc" +
	"tggggcaagccaggagggccagcatggggcagctgccaggggcgcagccgacaggcaggt" +
	"gttcggcgccagcctctccagctgccccaacaggtgcccaggcgctgggagggcggtgac" +
	"tcacgcgggccctgtgggagaaccagctttgcagacaggcgccaccagtgccccctcctc" +
	"tgcgatccaggagggacaactttgggttcttctgggtgtgtctccttcttttgtaggttc" +
	"tgcacccacccccacccccagccccaaagtctcggttcctatgagccgtgtgggtcagcc" +
	"accattcccgccaccccgggtccctgcgtcctttagttctcctggcccagggcctccaac" +
	"cttccagctgtcccacaaaaccccttcttgcaagggctttccagggcctggggccagggc" +
	"tggaaggaggatgcttccgcttctgccagctgccttgtctgcccacctcctccccaagcc" +
	"caggactcgggctcactggtcactggtttctttcattcccagcaccctgctcctctggcc" +
	"ctcatatgtctggccctcagtgactggtgtttggtttttggcctgtgtgtaacaaactgt" +
	"gtgtgacacttgtttcctgtttctccgccttcccctgcttcctcttgtgtccatctcttt" +
	"ctgacccaggcctggttcctttccctcctcctcccatttcacagatgggaaggtggcggc" +
	"caagaagggccaggccattcagcctctggaaaaaccttctcccaacctcccacagcccct" +
	"aatgactctcctggcctccctttagtagaggatgaagttgggttggcagggtaaactgag" +
	"accgggtggggtaggggtctggcgctcccgggaggagcactccttttgtggcccgagctg" +
	"catctcgcggcccctcccctgcaaggcctggggcgggggagggggccagggttcctgctg" +
	"ccttaaaagggctcaatgtcttggctctctcctccctcccccgtcctcagccctggctgg" +
	"ttcgtccctgctggcccactctcccggaaccccccggaacccctctctttcctccagaac" +
	"ccactgtctcctctccttccctcccctcccatacccatccctctctccatcctgcctcca" +
	"cttcttccacccccgggagtccaggcctccctgtccccacagtccctgagccacaagcct" +
	"ccaccccagctggtcccccacccaggctgcccagtttaacattcctagtcataggacctt" +
	"gacttctgagaggcctgattgtcatctgtaaataaggggtaggactaaagcactcctcct" +
	"ggaggactgagagatgggctggaccggagcacttgagtctgggatatgtgaccatgctac" +
	"ctttgtctccctgtcctgttccttcccccagccccaaatccagggttttccaaagtgtgg" +
	"ttcaagaaccacctgcatctgaatctagaggtactggatacaaccccacgtctgggccgt" +
	"tacccaggacattctacatgagaacgtgggggtggggccctggctgcacctgaactgtca" +
	"cctggagtcagggtggaaggtggaagaactgggtcttatttccttctccccttgttcttt" +
	"agggtctgtccttctgcagactccgttaccccaccctaaccatcctgcacacccttggag" +
	"ccctctgggccaatgccctgtcccgcaaagggcttctcaggcatctcacctctatgggag" +
	"ggcatttttggcccccagaaccttacacggtgtttatgtggggaagcccctgggaagcag" +
	"acagtcctagggtgaagctgagaggcagagagaaggggagacagacagagggtggggctt" +
	"tcccccttgtctccagtgccctttctggtgaccctcggttcttttcccccaccacccccc" +
	"cagcggagcccatcgtggtgaggcttaaggaggtccgactgcagagggacgacttcgaga" +
	"ttctgaaggtgatcggacgcggggcgttcagcgaggtaagccgaaccgggcgggagcctg" +
	"acttgactcgtggtgggcggggcataggggttggggcggggccttagaaattgatgaatg" +
	"accgagccttagaacctagggctgggctggaggcggggcttgggaccaatgggcgtggtg" +
	"tggcaggtggggcggggccacggctgggtgcagaagcgggtggagttgggtctgggcgag" +
	"cccttttgttttcccgccgtctccactctgtctcactatctcgacctcaggtagcggtag" +
	"tgaagatgaagcagacgggccaggtgtatgccatgaagatcatgaacaagtgggacatgc" +
	"tgaagaggggcgaggtgaggggctgggcggacgtggggggctttgaggatccgcgccccg" +
	"tctccggctgcagctcctccgggtgccctgcaggtgtcgtgcttccgtgaggagagggac" +
	"gtgttggtgaatggggaccggcggtggatcacgcagctgcacttcgccttccaggatgag" +
	"aactacctggtgagctccgggccggggggactaggaagagggacaagagcccgtgctgtc" +
	"actggacgaggaggtggggagaggaagctctaggattgggggtgctgcccggaaacgtct" +
	"gtgggaaagtctgtgtgcggtaagagggtgtgtcaggtggatgaggggccttccctatct" +
	"gagacggggatggtgtccttcactgcccgtttctggggtgatctgggggactcttataaa" +
	"gatgtctctgttgcggggggtctcttacctggaatgggataggtcttcaggaattctaac" +
	"ggggccactgcctagggaaggagtgtctgggacctattctctgggtgttgggtggcctct" +
	"gggttctctttcccagaacatctcagggggagtgaatctgcccagtgacatcccaggaaa" +
	"gtttttttgtttgtgtttttttttgaggggcgggggcgggggccgcaggtggtctctgat" +
	"ttggcccggcagatctctatggttatctctgggctggggctgcaggtctctgcccaagga" +
	"tggggtgtctctgggaggggttgtcccagccatccgtgatggatcagggcctcaggggac" +
	"taccaaccacccatgacgaaccccttctcagtacctggtcatggagtattacgtgggcgg" +
	"ggacctgctgacactgctgagcaagtttggggagcggattccggccgagatggcgcgctt" +
	"ctacctggcggagattgtcatggccatagactcggtgcaccggcttggctacgtgcacag" +
	"gtgggcgcagcatggccgaggggatagcaagcttgttccctggccgggttcttggaaggt" +
	"cagagcccagagaggccagggcctggagagggaccttcttggttggggcccaccgggggg" +
	"tgcctgggagtaggggtcagaactgtagaagccctacaggggcggaacccgaggaagtgg" +
	"ggtcccaggtggcactgcccggaggggcggagcctggtgggaccacagaagggaggttca" +
	"tttatcccacccttctcttttcctccgtgcagggacatcaaacccgacaacatcctgctg" +
	"gaccgctgtggccacatccgcctggccgacttcggctcttgcctcaagctgcgggcagat" +
	"ggaacggtgagccagtgccctggccacagagcaactggggctgctgatgagggatggaag" +
	"gcacagagtgtgggagcgggactggatttggaggggaaaagaggtggtgtgacccaggct" +
	"taagtgtgcatctgtgtggcggagtattagaccaggcagagggaggggctaagcatttgg" +
	"ggagtggttggaaggagggcccagagctggtgggcccagaggggtgggcccaagcctcgc" +
	"tctgctccttttggtccaggtgcggtcgctggtggctgtgggcaccccagactacctgtc" +
	"ccccgagatcctgcaggctgtgggcggtgggcctgggacaggcagctacgggcccgagtg" +
	"tgactggtgggcgctgggtgtattcgcctatgaaatgttctatgggcagacgcccttcta" +
	"cgcggattccacggcggagacctatggcaagatcgtccactacaaggtgagcacggccgc" +
	"agggagacctggcctctcccggtaggcgctcccagctatcgcctcctctccctctgagca" +
	"ggagcacctctctctgccgctggtggacgaaggggtccctgaggaggctcgagacttcat" +
	"tcagcggttgctgtgtcccccggagacacggctgggccggggtggagcaggcgacttccg" +
	"gacacatcccttcttctttggcctcgactgggatggtctccgggacagcgtgcccccctt" +
	"tacaccggatttcgaaggtgccaccgacacatgcaacttcgacttggtggaggacgggct" +
	"cactgccatggtgagcgggggcggggtaggtacctgtggcccctgctcggctgcgggaac" +
	"ctccccatgctccctccataaagttggagtaaggacagtgcctaccttctggggtcctga" +
	"atcactcattccccagagcacctgctctgtgcccatctactactgaggacccagcagtga" +
	"cctagacttacagtccagtgggggaacacagagcagtcttcagacagtaaggccccagag" +
	"tgatcagggctgagacaatggagtgcagggggtgggggactcctgactcagcaaggaagg" +
	"tcctggagggctttctggagtggggagctatctgagctgagacttggagggatgagaagc" +
	"aggagaggactcctcctcccttaggccgtctctcttcaccgtgtaacaagctgtcatggc" +
	"atgcttgctcggctctgggtgcccttttgctgaacaatactggggatccagcacggacca" +
	"gatgagctctggtccctgccctcatccagttgcagtctagagaattagagaattatggag" +
	"agtgtggcaggtgccctgaagggaagcaacaggatacaagaaaaaatgatggggccaggc" +
	"acggtgctcacgcctgtaaccccagcaatttggcaggccgaagtgggtggattgcttgag" +
	"cccaggagttcgagaccagcctgggcaatgtggtgagacccccgtctctacaaaaatgtt" +
	"ttaaaaattggttgggcgtggtggcgcatgcctgtatactcagctactagggtggccgac" +
	"gtgggcttgagcccaggaggtcaaggctgcagtgagctgtgattgtgccactgcactcca" +
	"gcctgggcaacggagagagactctgtctcaaaaataagataaactgaaattaaaaaatag" +
	"gctgggctggccgggcgtggtggctcacgcctgtaatctcagcactttgggaggccgagg" +
	"cgggtggatcacgaggtcagaagatggagaccagcctggccagcgtggcgaaaccccgtc" +
	"tctaccaaaaatataaaaaattagccaggcgtggtagagggcgcctgtaatctcagctac" +
	"tcaggacgctgaggcaggagaatcgcctgaacctgggaggcggaggttgcagtgagctga" +
	"gattgcaccactgcactccagcctgggtaacagagcgagactccgtatcaaagaaaaaga" +
	"aaaaagaaaaaatgctggaggggccactttagataagccctgagttggggctggtttggg" +
	"gggaacatgtaagccaagatcaaaaagcagtgaggggcccgccctgacgactgctgctca" +
	"catctgtgtgtcttgcgcaggagacactgtcggacattcgggaaggtgcgccgctagggg" +
	"tccacctgccttttgtgggctactcctactcctgcatggccctcaggtaagcactgccct" +
	"ggacggcctccaggggccacgaggctgcttgagcttcctgggtcctgctccttggcagcc" +
	"aatggagttgcaggatcagtcttggaaccttactgttttgggcccaaagactcctaagag" +
	"gccagagttggaggaccttaaattttcagatctatgtacttcaaaatgttagattgaatt" +
	"ttaaaacctcagagtcacagactgggcttcccagaatcttgtaaccattaacttttacgt" +
	"ctgtagtacacagagccacaggacttcagaacttggaaaatatgaagtttagacttttac" +
	"aatcagttgtaaaagaatgcaaattctttgaatcagccatataacaataaggccatttaa" +
	"aagtattaatttaggcgggccgcggtggctcacgcctgtaatcctagcactttgggaggc" +
	"caaggcaggtggatcatgaggtcaggagatcgagaccatcctggctaacacggtgaaacc" +
	"ccgtctctactaaaaatacaaaaaaattagccgggcatggtggcgggcgcttgcggtccc" +
	"agctacttgggaggcgaggcaggagaatggcatgaacccgggaggcggagcttgcagtga" +
	"gccgagatcatgccactgcactccagcctgggcgacagagcaagactccgtctcaaaaaa" +
	"aaaaaaaaaaaaagtatttatttaggccgggtgtggtggctcacgcctgtaattccagtg" +
	"ctttgggaggatgaggtgggtggatcacctgaggtcaggagttcgagaccagcctgacca" +
	"acgtggagaaacctcatctctactaaaaaacaaaattagccaggcatggtggcatatacc" +
	"tgtaatcccagctactcaggaggctgaggcaggagaatcagaacccaggagggggaggtt" +
	"gtggttagctgagatcgtgccattgcattccagcctgggcaacaagagtgaaacttcatc" +
	"tcaaaaaaaaaaaaaaaaaagtactaatttacaggctgggcatggtggctcacgcttgga" +
	"atcccagcactttgggaggctgaagtggacggattgcttcagcccaggagttcaagacca" +
	"gcctgagcaacataatgagaccctgtctctacaaaaaattgaaaaaatcgtgccaggcat" +
	"ggtggtctgtgcctgcagtcctagctactcaggagtctgaagtaggagaatcacttgagc" +
	"ctggagtttgaggcttcagtgagccatgatagattccagcctaggcaacaaagtgagacc" +
	"tggtctcaacaaaagtattaattacacaaataatgcattgcttatcacaagtaaattaga" +
	"aaatacagataaggaaaaggaagttgatatctcgtgagctcaccagatggcagtggtccc" +
	"tggctcacacgtgtactgacacatgtttaaatagtggagaacaggtgtttttttggtttg" +
	"tttttttccccttcctcatgctactttgtctaagagaacagttggttttctagtcagctt" +
	"ttattactggacaacattacacatactataccttatcattaatgaactccagcttgattc" +
	"tgaaccgctgcggggcctgaacggtgggtcaggattgaacccatcctctattagaaccca" +
	"ggcgcatgtccaggatagctaggtcctgagccgtgttcccacaggagggactgctgggtt" +
	"ggaggggacagccacttcataccccagggaggagctgtccccttcccacagctgagtggg" +
	"gtgtgctgacctcaagttgccatcttggggtcccatgcccagtcttaggaccacatctgt" +
	"ggaggtggccagagccaagcagtctccccatcaggtcggcctccctgtcctgaggccctg" +
	"agaagaggggtctgcagcggtcacatgtcaagggaggagatgagctgaccctagaacatg" +
	"ggggtctggaccccaagtccctgcagaaggtttagaaagagcagctcccaggggcccaag" +
	"gccaggagaggggcagggcttttcctaagcagaggaggggctattggcctacctgggact" +
	"ctgttctcttcgctctgctgctccccttcctcaaatcaggaggtcttggaagcagctgcc" +
	"cctacccacaggccagaagttctggttctccaccagataatcagcattctgtctccctcc" +
	"ccactccctcctcctctccccagggacagtgaggtcccaggccccacacccatggaagtg" +
	"gaggccgagcagctgcttgagccacacgtgcaagcgcccagcctggagccctcggtgtcc" +
	"ccacaggatgaaacagtaagttggtggaggggagggggtccgtcagggacaattgggaga" +
	"gaaaaggtgagggcttcccgggtggcgtgcactgtagagccctctagggacttcctgaac" +
	"agaagcagacagaaaccacggagagacgaggttacttcagacatgggacggtctctgtag" +
	"ttacagtggggcattaagtaagggtgtgtgtgttgctggggatctgagaagtcgatcttt" +
	"gagctgagcgctggtgaaggagaaacaagccatggaaggaaaggtgccaagtggtcaggc" +
	"gagagcctccagggcaaaggccttgggcaggtgggaatcctgatttgttcctgaaaggta" +
	"gtttggctgaatcattcctgagaaggctggagaggccagcaggaaacaaaacccagcaag" +
	"gccttttgtcgtgagggcattagggagctggagggattttgagcagcagagggacatagg" +
	"ttgtgttagtgtttgagcaccagccctctggtccctgtgtagatttagaggaccagactc" +
	"agggatggggctgagggaggtagggaagggagggggcttggatcattgcaggagctatgg" +
	"ggattccagaaatgttgaggggacggaggagtaggggataaacaaggattcctagcctgg" +
	"aaccagtgcccaagtcctgagtcttccaggagccacaggcagccttaagcctggtcccca" +
	"tacacaggctgaagtggcagttccagcggctgtccctgcggcagaggctgaggccgaggt" +
	"gacgctgcgggagctccaggaagccctggaggaggaggtgctcacccggcagagcctgag" +
	"ccgggagatggaggccatccgcacggacaaccagaacttcgccaggtcgggatcggggcc" +
	"ggggccggggccgggatgcgggccggtggcaacccttggcagcccctctcgtccggcccg" +
	"gacggactcaccgtccttacctccccacagtcaactacgcgaggcagaggctcggaaccg" +
	"ggacctagaggcacacgtccggcagttgcaggagcggatggagttgctgcaggcagaggg" +
	"agccacaggtgagtccctcatgtgtccccttccccggaggaccgggaggaggtgggccgt" +
	"ctgctccgcggggcgtgtatagacacctggaggagggaagggacccacgctggggcacgc" +
	"cgcgccaccgccctccttcgcccctccacgcgccctatgcctctttcttctccttccagc" +
	"tgtcacgggggtccccagtccccgggccacggatccaccttcccatgtaagacccctctc" +
	"tttcccctgcctcagacctgctgcccattctgcagatcccctccctggctcctggtctcc" +
	"ccgtccagatatagggctcaccctacgtctttgcgactttagagggcagaagccctttat" +
	"tcagccccagatctccctccgttcaggcctcaccagattccctccgggatctccctagat" +
	"aacctccccaacctcgattccgctcgctgtctctcgccccaccgctgagggctgggctgg" +
	"gctccgatcgggtcacctgtcccttctctctccagctagatggccccccggccgtggctg" +
	"tgggccagtgcccgctggtggggccaggccccatgcaccgccgccacctgctgctccctg" +
	"ccagggtacgtccggctgcccacgcccccctccgccgtcgcgccccgcgctccacccgcc" +
	"ccgtgccacccgcttagctgcgcatttgcggggctgggcccacggcaggagggcggatct" +
	"tcgggcagccaatcaacacaggccgctaggaagcagccaatgacgagttcggacgggatt" +
	"cgaggcgtgcgagtggactaacaacagctgtaggctgttggggcgggggcggggcgcagg" +
	"gaagagtgcgggcccacctatgggcgtaggcggggcgagtcccaggagccaatcagaggc" +
	"ccatgccgggtgttgacctcgccctctccccgcaggtccctaggcctggcctatcggagg" +
	"cgctttccctgctcctgttcgccgttgttctgtctcgtgccgccgccctgggctgcattg" +
	"ggttggtggcccacgccggccaactcaccgcagtctggcgccgcccaggagccgcccgcg" +
	"ctccctgaaccctagaactgtcttcgactccggggccccgttggaagactgagtgcccgg" +
	"ggcacggcacagaagccgcgcccaccgcctgccagttcacaaccgctccgagcgtgggtc" +
	"tccgcccagctccagtcctgtgatccgggcccgccccctagcggccggggagggaggggc" +
	"cgggtccgcggccggcgaacggggctcgaagggtccttgtagccgggaatgctgctgctg" +
	"ctgctgctgctgctgctgctgctggggggatcacagaccatttctttctttcggccaggc" +
	"tgaggccctgacgtggatgggcaaactgcaggcctgggaaggcagcaagccgggccgtcc" +
	"gtgttccatcctccacgcacccccacctatcgttggttcgcaaagtgcaaagctttcttg" +
	"tgcatgacgccctgctctggggagcgtctggcgcgatctctgcctgcttactcgggaaat" +
	"ttgcttttgccaaacccgctttttcggggatcccgcgcccccctcctcacttgcgctgct" +
	"ctcggagccccagccggctccgcccgcttcggcggtttggatatttattgacctcgtcct" +
	"ccgactcgctgacaggctacaggacccccaacaaccccaatccacgttttggatgcactg" +
	"agaccccgacattcctcggtatttattgtctgtccccacctaggacccccacccccgacc" +
	"ctcgcgaataaaaggccctccatctgcccaaagctctggactccacagtgtccgcggttt" +
	"gcgttgtgggccggagctccgcagcgggccaatccggaggcgtgtggaggcggccgaagg" +
	"tctgggaggagctagcgggatgcgaagcggccgaatcagggttgggggaggaaaagccac" +
	"ggggcggggctttggcgtccggccaataggagggcgagcgggccacccggaggcaccgcc" +
	"cccgcccagctgtggcccagctgtgccaccgagcgtcgagaagagggggctgggctggca" +
	"gcgcgcgcggccatcctccttccactgcgcctgcgcacgccacgcgcatccgctcctggg" +
	"acgcaagctcgagaaaagttgctgcaaactttctagcccgttccccgcccctcctcccgg" +
	"ccagacccgccccccctgcggagccgggaattc"
