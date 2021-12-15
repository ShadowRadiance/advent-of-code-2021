require './polymerizer'

RSpec.describe 'Polymerizer' do
  let(:polymerizer) { Polymerizer.new(rules) }

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

    it "calculates the next template" do
      expect(polymerizer.step(template)).to eq("")
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

    it "calculates the next template" do
      expect(polymerizer.step(template)).to eq("NCNBCHB")
    end

    context "after 3 steps" do
      before do
        @result = polymerizer.steps(3, template)
      end

      it "has the correct result" do
        expect(@result).to eq("NBBBCNCCNBBNBNBBCHBHHBCHB")
      end
    end

    context "after 4 steps" do
      before do
        @result = polymerizer.steps(4, template)
      end

      it "has the correct result" do
        expect(@result).to eq("NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB")
      end
    end

    context "after 10 steps" do
      before do
        @result = polymerizer.steps(10, template)
      end

      it "has most common element (B 1749) times" do
        counts = @result.chars.group_by(&:itself).transform_values(&:size) # {"B": 24, "C": 45, ...}
        max_letter, max_count = counts.max_by(&:last)

        expect(max_letter).to eq("B")
        expect(max_count).to eq(1749)
      end

      it "has least common element (H 161) times" do
        counts = @result.chars.group_by(&:itself).transform_values(&:size) # {"B": 24, "C": 45, ...}
        min_letter, min_count = counts.min_by(&:last)

        expect(min_letter).to eq("H")
        expect(min_count).to eq(161)
      end

      it "has variation 1588" do
        counts = @result.chars.group_by(&:itself).transform_values(&:size) # {"B": 24, "C": 45, ...}
        min_count, max_count = counts.values.minmax

        expect(max_count - min_count).to eq(1588)
      end
    end


  end
end