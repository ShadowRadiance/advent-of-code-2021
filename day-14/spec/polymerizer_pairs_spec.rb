require './polymerizer_pairs'

RSpec.describe 'PolymerizerPairs' do
  let(:polymerizer) { PolymerizerPairs.new(rules, template) }

  context "with an empty template" do
    let(:template) { "" }
    let(:rules) { {
      "CH" => "B",
      "HH" => "N",
      "CB" => "H",
      "NH" => "C",
      "HB" => "C",
      "HC" => "B",
      "HN" => "C",
      "NN" => "C",
      "BH" => "H",
      "NC" => "B",
      "NB" => "B",
      "BN" => "B",
      "BB" => "N",
      "BC" => "B",
      "CC" => "N",
      "CN" => "C",
    } }

    context "after one step" do
      before { polymerizer.step }
    end
    it "calculates the next template" do
      expect(polymerizer.pairs.values).to all(eq(0))
    end
  end

  context "with the example" do
    let(:template) { "NNCB" }
    let(:rules) { {
      "CH" => "B",
      "HH" => "N",
      "CB" => "H",
      "NH" => "C",
      "HB" => "C",
      "HC" => "B",
      "HN" => "C",
      "NN" => "C",
      "BH" => "H",
      "NC" => "B",
      "NB" => "B",
      "BN" => "B",
      "BB" => "N",
      "BC" => "B",
      "CC" => "N",
      "CN" => "C",
    } }

    context "after 1 steps" do
      before { polymerizer.steps(1) }

      it "has the correct counts" do
        # "NCNBCHB"
        expect(polymerizer.pairs["NC"]).to eq(1)
        expect(polymerizer.pairs["CN"]).to eq(1)
        expect(polymerizer.pairs["NB"]).to eq(1)
        expect(polymerizer.pairs["BC"]).to eq(1)
        expect(polymerizer.pairs["CH"]).to eq(1)
        expect(polymerizer.pairs["HB"]).to eq(1)
      end
    end

    context "after 3 steps" do
      before { polymerizer.steps(3) }

      it "has the correct counts" do
        expect(polymerizer.pairs).to eq({
          "CH" => 2,
          "HH" => 1,
          "CB" => 0,
          "NH" => 0,
          "HB" => 3,
          "HC" => 0,
          "HN" => 0,
          "NN" => 0,
          "BH" => 1,
          "NC" => 1,
          "NB" => 4,
          "BN" => 2,
          "BB" => 4,
          "BC" => 3,
          "CC" => 1,
          "CN" => 2,
        }) # eq("NBBBCNCCNBBNBNBBCHBHHBCHB")
      end
    end

    context "after 4 steps" do
      before { polymerizer.steps(4) }

      it "has the correct counts" do
        expect(polymerizer.pairs).to eq({
          "CH" => 0,
          "HH" => 1,
          "CB" => 5,
          "NH" => 1,
          "HB" => 0,
          "HC" => 3,
          "HN" => 1,
          "NN" => 0,
          "BH" => 3,
          "NC" => 1,
          "NB" => 9,
          "BN" => 6,
          "BB" => 9,
          "BC" => 4,
          "CC" => 2,
          "CN" => 3,
        }) # eq("NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB")
      end
    end

    context "after 10 steps" do
      before { polymerizer.steps(10) }

      it "has most common element (B 1749) times" do
        expect(polymerizer.most_common).to eq(["B", 1749])
      end

      it "has least common element (H 161) times" do
        expect(polymerizer.least_common).to eq(["H",161])
      end

      it "has variation 1588" do
        expect(polymerizer.variance).to eq(1588)
      end
    end
  end
end