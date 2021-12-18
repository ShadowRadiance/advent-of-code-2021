require "./heap"

RSpec.describe Heap do
  context "with an empty heap" do
    subject { Heap.new }
    
    it { is_expected.to be_empty }

    context "after adding one item" do
      before do
        subject.add(54)
      end
      it { is_expected.not_to be_empty }
      it { is_expected.to have_attributes(size: 1) }
      it { is_expected.to have_attributes(to_a: [54]) }
    end

    context "after adding four items" do
      before do
        subject.add(54)
        subject.add(2)
        subject.add(12)
        subject.add(1)
      end
      it { is_expected.not_to be_empty }
      it { is_expected.to have_attributes(size: 4) }
      it { is_expected.to have_attributes(to_a: [1, 2, 12, 54]) }

      context "after removing the top item" do
        before do
          @top = subject.extract
        end
        
        it "extracted the lowest value" do
          expect(@top).to eq(1)
        end

        it { is_expected.not_to be_empty }
        it { is_expected.to have_attributes(size: 3) }
        it { is_expected.to have_attributes(to_a: [2, 54, 12]) }
      end
    end
  end
end