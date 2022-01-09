# frozen_string_literal: true

RSpec.describe Stack do
  subject { described_class.new(capacity: capacity) }
  let(:capacity) { 2 }

  context "initial state" do
    it "is empty" do
      expect(subject).to be_empty
      expect(subject.size).to eq(0)
      expect(subject.empty?).to eq(true)
    end

    it "is not full" do
      expect(subject).not_to be_full
      expect(subject.full?).to eq(false)
    end

    it "raises errors on read" do
      expect { subject.top }.to raise_error(Stack::Underflow)
      expect { subject.pop }.to raise_error(Stack::Underflow)
    end

    it "does not raise errors on write" do
      expect { subject.push(0) }.not_to raise_error
    end

    it "increases the size of the stack on write" do
      expect { subject.push(0) }.to change { subject.size }.from(0).to(1)
    end
  end

  context "with a single element" do
    before { subject.push("check") }

    it "is not empty" do
      expect(subject).not_to be_empty
      expect(subject.size).to eq(1)
      expect(subject.empty?).to eq(false)
    end

    it "is not full" do
      expect(subject).not_to be_full
      expect(subject.full?).to eq(false)
    end

    it "returns the pushed value(s)" do
      expect(subject.top).to eq("check")
      expect(subject.pop).to eq("check")
    end

    it "does not raise errors on top/pop" do
      expect { subject.top }.not_to raise_error
      expect { subject.pop }.not_to raise_error
    end

    it "does not raise errors on write" do
      expect { subject.push(0) }.not_to raise_error
    end

    it "increases the size of the stack on write" do
      expect { subject.push(0) }.to change { subject.size }.from(1).to(2)
    end
  end

  context "with capacity elements" do
    before { subject.push("check"); subject.push("second") }

    it "is not empty" do
      expect(subject).not_to be_empty
      expect(subject.size).to eq(2)
      expect(subject.empty?).to eq(false)
    end

    it "is not full" do
      expect(subject).to be_full
      expect(subject.full?).to eq(true)
    end

    it "returns the pushed value(s)" do
      expect(subject.top).to eq("second")
      expect(subject.pop).to eq("second")

      expect(subject.top).to eq("check")
      expect(subject.pop).to eq("check")
    end

    it "does not raise errors on top/pop" do
      expect { subject.top }.not_to raise_error
      expect { subject.pop }.not_to raise_error
    end

    it "does not raise errors on write" do
      expect { subject.push(0) }.to raise_error(Stack::Overflow)
    end
  end
end