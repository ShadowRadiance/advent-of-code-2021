require "./bingo/card"

RSpec.describe 'Bingo Card' do
  subject { Bingo::Card.new(data) }

  context 'with no data' do
    let(:data) { [] }

    it "raises an error" do
      expect { subject }.to raise_error(ArgumentError)
    end
  end

  context 'with the sample data' do
    let(:data) {
      [
        14, 21, 17, 24,  4,
        10, 16, 15,  9, 19,
        18,  8, 23, 26, 20,
        22, 11, 13,  6,  5,
         2,  0, 12,  3,  7,
      ]
    }

    context 'initial state' do
      it "has no bingo" do
        expect(subject.bingo?).to be(false)
      end

      it "shows the marked numbers as marked" do
        expect(subject.marked).to eq(Set[])
      end

      it "shows the unmarked numbers as unmarked" do
        expect(subject.unmarked).to eq(Set[ 14, 21, 17, 24,  4, 
                                            10, 16, 15,  9, 19, 
                                            18,  8, 23, 26, 20, 
                                            22, 11, 13,  6,  5, 
                                             2,  0, 12,  3,  7,])
      end

      it "can separate the lines" do
        expect(subject.lines).to eq([
          [14, 21, 17, 24,  4], 
          [10, 16, 15,  9, 19], 
          [18,  8, 23, 26, 20], 
          [22, 11, 13,  6,  5], 
          [ 2,  0, 12,  3,  7],
        ])
      end

      it "can separate the columns" do
        expect(subject.columns).to eq([
          [14, 10, 18, 22,  2], 
          [21, 16,  8, 11,  0], 
          [17, 15, 23, 13, 12], 
          [24,  9, 26,  6,  3], 
          [ 4, 19, 20,  5,  7]
        ])
      end
    end

    context 'after calling a number' do
      before do
        [7].each { |number| subject.mark(number)}
      end

      it "has no bingo" do
        expect(subject.bingo?).to be(false)
      end
  
      it "shows the marked numbers as marked" do
        expect(subject.marked).to eq(Set[7])
      end

      it "shows the unmarked numbers as unmarked" do
        expect(subject.unmarked).to eq(Set[ 14, 21, 17, 24,  4, 
                                            10, 16, 15,  9, 19, 
                                            18,  8, 23, 26, 20, 
                                            22, 11, 13,  6,  5, 
                                             2,  0, 12,  3,    ])
      end
    end

    context "after calling enough numbers for a bingo column" do
      before do
        [4,7,19,20,5].each { |number| subject.mark(number)}
      end

      it "has a bingo" do
        expect(subject.bingo?).to be(true)
      end
  
      it "shows the marked numbers as marked" do
        expect(subject.marked).to eq(Set[4,7,19,20,5])
      end

      it "shows the unmarked numbers as unmarked" do
        expect(subject.unmarked).to eq(Set[ 14, 21, 17, 24,
                                            10, 16, 15,  9,
                                            18,  8, 23, 26,
                                            22, 11, 13,  6,
                                             2,  0, 12,  3, ])
    end

    end

    context "after calling enough numbers for a bingo line" do
      before do
        [7,4,9,5,11,17,23,2,0,14,21,24].each { |number| subject.mark(number)}
      end

      it "has a bingo" do
        expect(subject.bingo?).to be(true)
      end
  
      it "shows the marked numbers as marked" do
        expect(subject.marked).to eq(Set[7,4,9,5,11,17,23,2,0,14,21,24])
      end

      it "shows the unmarked numbers as unmarked" do
        expect(subject.unmarked).to eq(Set[                    
                                            10, 16, 15,     19, 
                                            18,  8,     26, 20, 
                                            22,     13,  6,    
                                                    12,  3,    ])
      end
    end
  end
end