require './bingo/game'

RSpec.describe "Bingo Game" do
  subject { Bingo::Game.new(cards, numbers) }

  context "with no data" do
    let(:cards) { [] }
    let(:numbers) { }

    it "raises an error" do
      expect { subject }.to raise_error(ArgumentError)
    end
  end

  context "with the sample data" do
    let(:cards) { [
      [ 22, 13, 17, 11,  0,
        8,  2, 23,  4, 24,
       21,  9, 14, 16,  7,
        6, 10,  3, 18,  5,
        1, 12, 20, 15, 19, ],
     [  3, 15,  0,  2, 22,
        9, 18, 13, 17,  5,
       19,  8,  7, 25, 23,
       20, 11, 10, 24,  4,
       14, 21, 16, 12,  6, ],
     [ 14, 21, 17, 24,  4,
       10, 16, 15,  9, 19,
       18,  8, 23, 26, 20,
       22, 11, 13,  6,  5,
        2,  0, 12,  3,  7, ],
     ] }
    let(:numbers) { [ 7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1 ] }

    describe "at the start of the game" do
      it "does not have a winner" do
        expect(subject.winner).to be_nil
      end
    end

    describe "after one number" do
      it "does not have a winner" do
        subject.call_a_number
        expect(subject.winner).to be_nil
      end
    end

    describe "after twelve numbers" do
      before do
        12.times { 
          subject.call_a_number
          subject.determine_winner
        }
      end

      it "has a winner" do
        expect(subject.winner).not_to be_nil
      end
      
      it "has a winning score of 4512" do
        expect(subject.winning_score).to eq(4512)
      end
    end

    describe "running the game until bingo" do
      before do
        subject.play
      end

      it "has a winner" do
        expect(subject.winner).not_to be_nil
      end
      
      it "has a winning score of 4512" do
        expect(subject.winning_score).to eq(4512)
      end

      it "took twelve turns" do
        expect(subject.turns).to eq(12)
      end
    end

    describe "running the game ranking all boards" do
      before do
        subject.play_all_cards
      end

      it "ranks all the boards" do
        expect(subject.ranked_cards).not_to be_empty
      end

      it "has a losing score of 1924" do
        expect(subject.ranked_cards.last[:score]).to eq(1924)
      end

      it "took fifteen turns" do
        expect(subject.turns).to eq(15)
      end
    end
  end

end